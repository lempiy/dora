package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lempiy/dora/services/api/handlers"
	"github.com/lempiy/dora/services/api/clients"
)

const port = "9000"

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	r := e.Router()

	botService := clients.NewBotClient()
	handlers.Run(r, botService)

	e.Logger.Fatal(e.Start(":" + port))
}
