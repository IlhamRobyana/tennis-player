package route

import (
	"github.com/labstack/echo"

	"github.com/ilhamrobyana/tennis-player/container"
)

func containerRoute(e *echo.Echo) {
	g := e.Group("/container")
	g.POST("/", container.Create)
	g.PUT("/put-ball", container.PutBall)
}
