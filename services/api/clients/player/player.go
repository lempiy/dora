package player

import (
	"github.com/labstack/echo"
	"github.com/lempiy/dora/shared/pb/bot"
	"github.com/lempiy/dora/services/api/defs"
	"net/http"
)

func GetMatchesHistory(botService bot.BotServiceClient) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var data defs.PlayerIdRequest
		if err := ctx.Bind(&data); err != nil {
			return ctx.String(http.StatusInternalServerError, "Server error")
		}
		if data.PlayerID == 0 {

		}
	}
}