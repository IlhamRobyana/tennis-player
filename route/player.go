package route

import (
	"github.com/labstack/echo"

	"github.com/ilhamrobyana/tennis-player/mwcustom"
	"github.com/ilhamrobyana/tennis-player/player"
)

func playerRoute(e *echo.Echo) {
	g := e.Group("/player")
	g.POST("/signup", player.Signup)
	g.POST("/login", player.Login)
	g.GET("/play", player.Play, mwcustom.Authorization)
}
