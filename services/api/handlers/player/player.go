package player

import (
	"github.com/labstack/echo"
	"github.com/lempiy/dora/shared/pb/bot"
	"github.com/lempiy/dora/services/api/defs"
	"net/http"
	"golang.org/x/net/context"
)

func GetMatchesHistory(botService bot.BotServiceClient) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var data defs.PlayerIdRequest
		if err := ctx.Bind(&data); err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		if data.PlayerID == 0 {
			return ctx.JSON(http.StatusForbidden, defs.StatusResponse{
				Status: "fail",
				Info: "Player ID cannot be empty",
			})
		}
		req := bot.MatchesHistoryRequest{
			PlayerId: data.PlayerID,
		}
		result, err := botService.GetMatchesHistory(context.Background(), &req)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, defs.StatusResponse{
				Status: "fail",
				Info: err.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, result.Matches)
	}
}

func GetMatchDetails(botService bot.BotServiceClient) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var data defs.MatchIdRequest
		if err := ctx.Bind(&data); err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		if data.MatchID == 0 {
			return ctx.JSON(http.StatusForbidden, defs.StatusResponse{
				Status: "fail",
				Info: "Match ID cannot be empty",
			})
		}
		req := bot.MatchDetailsRequest{
			MatchId: data.MatchID,
		}
		result, err := botService.GetMatchDetails(context.Background(), &req)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, defs.StatusResponse{
				Status: "fail",
				Info: err.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, result.Match)
	}
}

func GetPlayerCard(botService bot.BotServiceClient) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var data defs.PlayerIdRequest
		if err := ctx.Bind(&data); err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		if data.PlayerID == 0 {
			return ctx.JSON(http.StatusForbidden, defs.StatusResponse{
				Status: "fail",
				Info: "Player ID cannot be empty",
			})
		}
		req := bot.PlayerCardRequest{
			PlayerId: data.PlayerID,
		}
		result, err := botService.GetPlayerCard(context.Background(), &req)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, defs.StatusResponse{
				Status: "fail",
				Info: err.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, result)
	}
}
