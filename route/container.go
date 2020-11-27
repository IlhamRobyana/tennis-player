package route

import (
	"github.com/labstack/echo"

	"github.com/ilhamrobyana/tennis-player/container"
	"github.com/ilhamrobyana/tennis-player/mwcustom"
)

func containerRoute(e *echo.Echo) {
	g := e.Group("/container")
	g.Use(mwcustom.Authorization)
	g.POST("/", container.Create)
	g.PUT("/put-ball", container.PutBall)
}
