package handlers

import (
	"github.com/labstack/echo"
	"github.com/lempiy/dora/shared/pb/bot"
	"net/http"
	"github.com/lempiy/dora/services/api/handlers/player"
	"github.com/lempiy/dora/shared/pb/prs"
	"github.com/lempiy/dora/services/api/handlers/replay"
)

func Run(r *echo.Router, botService bot.BotServiceClient, parseService prs.ParseServiceClient) {
	r.Add(http.MethodPost, "/history", player.GetMatchesHistory(botService))
	r.Add(http.MethodPost, "/card", player.GetPlayerCard(botService))
	r.Add(http.MethodPost, "/match", player.GetMatchDetails(botService))

	r.Add(http.MethodPost, "/parse", replay.ParseReplay(parseService))
}
