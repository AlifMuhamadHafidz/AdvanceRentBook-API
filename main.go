package main

import (
	"advancerentbook-api/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	config.Migrate(db)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
}
