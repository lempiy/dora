package service

import (
	"fmt"
	"github.com/lempiy/dora/services/parser/utils"
	"github.com/lempiy/dora/shared"
	"golang.org/x/net/context"
	"log"
	"os"
)

const tempLocation = "./.temp"

type ParserService struct{}

func (s *ParserService) Parse(ctx context.Context, req *shared.ParseRequest) (*shared.ParseResult, error) {
	path := fmt.Sprintf("%s/%d_%d.dem.bz2", tempLocation, req.MatchId, req.ReplaySalt)
	err := utils.DownloadFile(path, req.ReplayUrl)
	if err != nil {
		log.Printf("ParserService.Parse: %s", err)
		return &shared.ParseResult{
			ReplayData: nil,
			Success:    false,
			ErrorInfo:  "Error upon downloading replay",
		}, nil
	}
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Printf("ParserService.Parse: %s", err)
		return &shared.ParseResult{
			ReplayData: nil,
			Success:    false,
			ErrorInfo:  "Error upon opening replay",
		}, nil
	}
	err = utils.UncompressBZ2(f, tempLocation)
	if err != nil {
		log.Printf("ParserService.Parse: %s", err)
		return &shared.ParseResult{
			ReplayData: nil,
			Success:    false,
			ErrorInfo:  "Error upon decompressing replay archive",
		}, nil
	}
	replayFile, err := os.Open(fmt.Sprintf("%s/%d_%d.dem", t
	if err != nil {
		log.Printf("ParserService.Parse: %s", err)
		return &shared.ParseResult{
			ReplayData: nil,
			Success:    false,
			ErrorInfo:  "Error upon decompressing replay archive",
		}, nil
	}
	defer replayFile.Close()
}
