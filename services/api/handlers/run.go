package handlers

import (
	"github.com/labstack/echo"
	"github.com/lempiy/dora/shared/pb/bot"
	"net/http"
	"github.com/lempiy/dora/services/api/handlers/player"
)

func Run(r *echo.Router, botService bot.BotServiceClient) {
	r.Add(http.MethodPost, "/history", player.GetMatchesHistory(botService))
	r.Add(http.MethodPost, "/card", player.GetPlayerCard(botService))
	r.Add(http.MethodPost, "/match", player.GetMatchDetails(botService))
}
