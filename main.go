package main

import (
	"advancerentbook-api/config"
	usrData "advancerentbook-api/features/user/data"
	usrHdl "advancerentbook-api/features/user/handler"
	usrSrv "advancerentbook-api/features/user/services"

	bookData "advancerentbook-api/features/book/data"
	bookHdl "advancerentbook-api/features/book/handler"
	bookSrv "advancerentbook-api/features/book/services"
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

	bData := bookData.New(db)
	bSrv := bookSrv.New(bData)
	bHdl := bookHdl.New(bSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	//users
	e.POST("/register", uHdl.Register())
	e.POST("/login", uHdl.Login())
	e.GET("/users", uHdl.Profile(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/users", uHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/users", uHdl.Deactivate(), middleware.JWT([]byte(config.JWTKey)))

	//books
	e.POST("/books", bHdl.Add(), middleware.JWT([]byte(config.JWTKey)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
