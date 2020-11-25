package route

import "github.com/labstack/echo"

func Init(e *echo.Echo) {
	playerRoute(e)
	containerRoute(e)
}
