package player

import (
	"net/http"

	"github.com/hipeid/backend/errcode"
	"github.com/ilhamrobyana/tennis-player/entity"
	pg "github.com/ilhamrobyana/tennis-player/pg_storage"
	"github.com/labstack/echo"
)

var coreInstance *core

func Signup(c echo.Context) (e error) {
	r := new(entity.SignupRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	player := entity.Player{0, r.Username, r.Password}
	playerCore := getCore()
	response, e := playerCore.signup(player)
	if e != nil {
		httpStatus := http.StatusInternalServerError
		if e.Error() == errcode.UserExists {
			httpStatus = http.StatusBadRequest
		}
		return c.JSON(httpStatus, map[string]interface{}{"message": e.Error()})
	}
	return c.JSON(http.StatusCreated, response)
}

func Login(c echo.Context) (e error) {
	r := new(entity.LoginRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	authCore := getCore()
	response, err := authCore.login(r.Username, r.Password)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}
	return c.JSON(http.StatusOK, response)
}

func Play(c echo.Context) (e error) {
	playerID := c.Get("id").(uint64)
	playerCore := getCore()
	container, err := playerCore.play(playerID)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		if container.ID == 0 {
			httpStatus := http.StatusOK
			return c.JSON(httpStatus, map[string]interface{}{"message": "No Container is filled yet"})
		}
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}
	response := entity.PlayResponse{"You played tennis with the following container, the balls are resetted to 0", container}
	return c.JSON(http.StatusOK, response)
}

func getCore() (c *core) {
	c = coreInstance

	if c == nil {
		c = new(core)
		c.playerStore = pg.Player{}
		c.containerStore = pg.Container{}
		coreInstance = c
	}

	return
}
