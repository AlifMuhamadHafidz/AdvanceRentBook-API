package main

import (
	"advancerentbook-api/config"
	usrData "advancerentbook-api/features/user/data"
	usrHdl "advancerentbook-api/features/user/handler"
	usrSrv "advancerentbook-api/features/user/services"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	config.Migrate(db)

	uData := usrData.New(db)
	uSrv := usrSrv.New(uData)
	uHdl := usrHdl.New(uSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	//users
	e.POST("/register", uHdl.Register())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
