package exec

import (
	"github.com/dotabuff/manta"
	"io"
	"time"
)

type Processor interface {
	Process(startGameTime, gameTime time.Duration, entity *manta.Entity, op manta.EntityOp) error
	Finish(gameEndTime time.Duration) error
}

type Parser struct {
	gameTime time.Duration
	preGameTime time.Duration
	gameEndTime time.Duration
	startGameTime time.Duration
	processors []Processor

	gameTotalTimeSec uint64
}

func NewParser() *Parser {
	return &Parser{}
}

func (parser *Parser) RegisterProcessors(processors ...Processor) {
	parser.processors = append(parser.processors, processors...)
}

// ParseReplay - parses replay from source stream, returns parsed data
func (parser *Parser) Parse(source io.Reader) error {
	p, err := manta.NewStreamParser(source)
	if err != nil {
		return err
	}
	p.OnEntity(func(entity *manta.Entity, op manta.EntityOp) error {
		if entity.GetClassName() == "CDOTAGamerulesProxy" {
			if v, ok := entity.GetFloat32("m_pGameRules.m_fGameTime"); ok {
				parser.gameTime = time.Duration(v)
			}
			if v, ok := entity.GetFloat32("m_pGameRules.m_flPreGameStartTime"); ok {
				parser.preGameTime = time.Duration(v)
			}
			if v, ok := entity.GetFloat32("m_pGameRules.m_flGameStartTime"); ok {
				parser.startGameTime = time.Duration(v)
			}
			if v, ok := entity.GetFloat32("m_pGameRules.m_flGameEndTime"); ok {
				parser.gameEndTime = time.Duration(v)
			}
		}
		for _, proc := range parser.processors {
			err = proc.Process(parser.startGameTime, parser.gameTime, entity, op)
			if err != nil {
				return err
			}
		}
		return nil
	})
	parser.gameTotalTimeSec = uint64(parser.gameEndTime - parser.startGameTime)
	p.Start()
	for _, proc := range parser.processors {
		err = proc.Finish(parser.gameEndTime)
		if err != nil {
			return err
		}
	}
	return nil
}

func (parser *Parser) GetGameTimeDuration() uint64 {
	return parser.gameTotalTimeSec
}
