package service

import (
	"fmt"
	"github.com/lempiy/dora/services/parser/exec"
	"github.com/lempiy/dora/services/parser/utils"
	"github.com/lempiy/dora/shared"
	"golang.org/x/net/context"
	"log"
	"os"
	"path/filepath"
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
	defer clearTempFolder()
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
	replayFile, err := os.Open(fmt.Sprintf("%s/%d_%d.dem", tempLocation, req.MatchId, req.ReplaySalt))
	if err != nil {
		log.Printf("ParserService.Parse: %s", err)
		return &shared.ParseResult{
			ReplayData: nil,
			Success:    false,
			ErrorInfo:  "Error upon decompressing replay archive",
		}, nil
	}
	defer replayFile.Close()
	data, err := exec.ParseReplay(replayFile)
	if err != nil {
		log.Printf("ParserService.Parse: %s", err)
		return &shared.ParseResult{
			ReplayData: nil,
			Success:    false,
			ErrorInfo:  "Error upon parsing replay",
		}, nil
	}
	return &shared.ParseResult{
		ReplayData: data,
		Success:    true,
	}, nil
}

func clearTempFolder() {
	d, err := os.Open(tempLocation)
	if err != nil {
		log.Printf("clearTempFolder: %s", err)
		return
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		log.Printf("clearTempFolder: %s", err)
		return
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(tempLocation, name))
		if err != nil {
			log.Printf("clearTempFolder: %s", err)
		}
	}
	return
}
