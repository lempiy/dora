package service

import (
	"fmt"
	"github.com/lempiy/dora/services/parser/exec"
	"github.com/lempiy/dora/services/parser/utils"
	"github.com/lempiy/dora/shared/pb/prs"
	"golang.org/x/net/context"
	"log"
	"os"
)


type ParserService struct{}

func (s *ParserService) Parse(ctx context.Context, req *prs.ParseRequest) (*prs.ParseResult, error) {
	tempLocation := createTempFolder(req.MatchId)
	defer clearTempFolder(tempLocation)
	path := fmt.Sprintf("%s/%d_%d.dem.bz2", tempLocation, req.MatchId, req.ReplaySalt)
	err := utils.DownloadFile(path, req.ReplayUrl)
	if err != nil {
		log.Printf("ParserService.Parse: Error upon DownloadFile: %s", err)
		return &prs.ParseResult{
			ReplayData: nil,
			Success:    false,
			ErrorInfo:  "Error upon downloading replay",
		}, nil
	}
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Printf("ParserService.Parse: Error upon os.Open: %s", err)
		return &prs.ParseResult{
			ReplayData: nil,
			Success:    false,
			ErrorInfo:  "Error upon opening replay",
		}, nil
	}
	err = utils.UncompressBZ2(f, fmt.Sprintf("%s/%d_%d.dem", tempLocation, req.MatchId, req.ReplaySalt))
	if err != nil {
		log.Printf("ParserService.Parse: Error upon UncompressBZ2: %s", err)
		return &prs.ParseResult{
			ReplayData: nil,
			Success:    false,
			ErrorInfo:  "Error upon decompressing replay archive",
		}, nil
	}
	replayFile, err := os.Open(fmt.Sprintf("%s/%d_%d.dem", tempLocation, req.MatchId, req.ReplaySalt))
	if err != nil {
		log.Printf("ParserService.Parse: Error upon os.Open: %s", err)
		return &prs.ParseResult{
			ReplayData: nil,
			Success:    false,
			ErrorInfo:  "Error upon decompressing replay archive",
		}, nil
	}
	defer replayFile.Close()
	data, err := exec.ParseReplay(replayFile)
	if err != nil {
		log.Printf("ParserService.Parse: Error upon exec.ParseReplay: %s", err)
		return &prs.ParseResult{
			ReplayData: nil,
			Success:    false,
			ErrorInfo:  "Error upon parsing replay",
		}, nil
	}
	return &prs.ParseResult{
		ReplayData: data,
		Success:    true,
	}, nil
}

func createTempFolder(matchId uint64) string {
	folder := fmt.Sprintf("./match_%d", matchId)
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		log.Printf("createTempFolder. Err: `%s`", err)
	}
	return folder
}


func clearTempFolder(folder string) {
	err := os.RemoveAll(folder)
	if err != nil {
		log.Printf("clearTempFolder. Err: `%s`", err)
	}
}
