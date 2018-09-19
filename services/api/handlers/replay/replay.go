package replay

import (
	"github.com/lempiy/dora/shared/pb/prs"
	"github.com/labstack/echo"
	"github.com/lempiy/dora/services/api/defs"
	"net/http"
	"golang.org/x/net/context"
)

func ParseReplay(parseService prs.ParseServiceClient) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var data defs.ParseReplayRequest
		if err := ctx.Bind(&data); err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		if data.PlayerId == 0 {
			return ctx.JSON(http.StatusForbidden, defs.StatusResponse{
				Status: "fail",
				Info: "Player ID cannot be empty",
			})
		}
		if data.MatchId == 0 {
			return ctx.JSON(http.StatusForbidden, defs.StatusResponse{
				Status: "fail",
				Info: "Match ID cannot be empty",
			})
		}
		if data.ReplaySalt == 0 {
			return ctx.JSON(http.StatusForbidden, defs.StatusResponse{
				Status: "fail",
				Info: "Replay Salt cannot be empty",
			})
		}
		if len(data.ReplayUrl) == 0 {
			return ctx.JSON(http.StatusForbidden, defs.StatusResponse{
				Status: "fail",
				Info: "Replay URL cannot be empty",
			})
		}
		req := prs.ParseRequest{
			ReplayUrl: data.ReplayUrl,
			ReplaySalt: data.ReplaySalt,
			PlayerId: data.PlayerId,
			MatchId: data.MatchId,
		}
		result, err := parseService.Parse(context.Background(), &req)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, defs.StatusResponse{
				Status: "fail",
				Info: err.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, result)
	}
}
