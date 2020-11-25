package main

import (
	"fmt"
	"os"

	"github.com/ilhamrobyana/tennis-player/config"
	"github.com/ilhamrobyana/tennis-player/route"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("can't load .env : %v", err))
	}

	port := os.Getenv("GOLANG_PORT")

	e := echo.New()
	e.Use(middleware.CORS())
	config.Init(e)
	route.Init(e)

	e.Logger.Fatal(e.Start(":" + port))
}
