package exec

import (
	"github.com/dotabuff/manta"
	"github.com/lempiy/dora/shared/pb/prs"
	"io"
	"strings"
	"time"
)

// ParseReplay - parses replay from source stream, returns parsed data
func ParseReplay(source io.Reader) (*prs.ReplayData, error) {
	var result prs.ReplayData
	p, err := manta.NewStreamParser(source)
	if err != nil {
		return nil, err
	}
	var (
		gameTime      = time.Duration(0)
		preGameTime   = time.Duration(0)
		startGameTime = time.Duration(0)
		gameEndTime   = time.Duration(0)
		movement      = make(map[string]*[]*prs.Move)
		checks        = make(map[string]time.Duration)
	)
	p.OnEntity(func(entity *manta.Entity, op manta.EntityOp) error {
		if entity.GetClassName() == "CDOTAGamerulesProxy" {
			if v, ok := entity.GetFloat32("m_pGameRules.m_fGameTime"); ok {
				gameTime = time.Duration(v)
			}
			if v, ok := entity.GetFloat32("m_pGameRules.m_flPreGameStartTime"); ok {
				preGameTime = time.Duration(v)
			}
			if v, ok := entity.GetFloat32("m_pGameRules.m_flGameStartTime"); ok {
				startGameTime = time.Duration(v)
			}
			if v, ok := entity.GetFloat32("m_pGameRules.m_flGameEndTime"); ok {
				gameEndTime = time.Duration(v)
			}
		}
		if strings.HasPrefix(entity.GetClassName(), "CDOTA_Unit_Hero") {
			if startGameTime != 0 && checks[entity.GetClassName()] != gameTime {
				checks[entity.GetClassName()] = gameTime
				x, _ := entity.GetUint64("CBodyComponent.m_cellX")
				y, _ := entity.GetUint64("CBodyComponent.m_cellY")
				if arr, exist := movement[entity.GetClassName()]; !exist {
					data := &[]*prs.Move{
						{
							Time: uint64(gameTime - startGameTime),
							X:    x,
							Y:    y,
						},
					}
					movement[entity.GetClassName()] = data
				} else {
					*arr = append(*arr, &prs.Move{
						Time: uint64(gameTime - startGameTime),
						X:    x,
						Y:    y,
					})
				}
			}

		}
		return nil
	})
	result.GameTotalTimeSec = uint64(gameEndTime - startGameTime)
	p.Start()
	for key, value := range movement {
		result.MovesMap = append(result.MovesMap, &prs.MovesMap{
			HeroName: key,
			Moves:    *value,
		})
	}
	return &result, nil
}
