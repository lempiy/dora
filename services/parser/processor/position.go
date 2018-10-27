package processor

import (
	"github.com/dotabuff/manta"
	"github.com/lempiy/dora/shared/pb/prs"
	"strings"
	"time"
)

const DefaultNetOffset = 50

type PositionProcessor struct{
	movement map[string]*[]*prs.Move
	checks map[string]time.Duration
	result []*prs.MovesMap
}

func NewPositionProcessor() *PositionProcessor {
	return &PositionProcessor{
		movement: make(map[string]*[]*prs.Move),
		checks: make(map[string]time.Duration),
	}
}

func (p *PositionProcessor) Process(startGameTime, gameTime time.Duration, entity *manta.Entity, op manta.EntityOp) error {
	if strings.HasPrefix(entity.GetClassName(), "CDOTA_Unit_Hero") {
		if startGameTime != 0 && p.checks[entity.GetClassName()] != gameTime {
			p.checks[entity.GetClassName()] = gameTime
			x, _ := entity.GetUint64("CBodyComponent.m_cellX")
			y, _ := entity.GetUint64("CBodyComponent.m_cellY")
			if arr, exist := p.movement[entity.GetClassName()]; !exist {
				data := &[]*prs.Move{
					{
						Time: uint64(gameTime - startGameTime),
						X:    x - DefaultNetOffset,
						Y:    y - DefaultNetOffset,
					},
				}
				p.movement[entity.GetClassName()] = data
			} else {
				*arr = append(*arr, &prs.Move{
					Time: uint64(gameTime - startGameTime),
					X:    x - DefaultNetOffset,
					Y:    y - DefaultNetOffset,
				})
			}
		}
	}
	return nil
}

func (p *PositionProcessor) Finish(gameEndTime time.Duration) error {
	for key, value := range p.movement {
		p.result = append(p.result, &prs.MovesMap{
			HeroName: key,
			Moves:    *value,
		})
	}
	return nil
}

func (p *PositionProcessor) Result() []*prs.MovesMap {
	return p.result
}
