package container

import (
	"net/http"

	"github.com/ilhamrobyana/tennis-player/entity"
	pg "github.com/ilhamrobyana/tennis-player/pg_storage"
	"github.com/labstack/echo"
)

var coreInstance *core

func Create(c echo.Context) (e error) {
	r := new(entity.ContainerCreateRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	container := new(entity.Container)
	container.Capacity = r.Capacity
	container.PlayerID = c.Get("id").(uint64)

	containerCore := getCore()
	createdContainer, err := containerCore.create(*container)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}
	return c.JSON(http.StatusCreated, createdContainer)
}

func PutBall(c echo.Context) (e error) {
	playerID := c.Get("id").(uint64)
	containerCore := getCore()
	updatedID, err := containerCore.putBall(playerID)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	} else if updatedID == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "all containers are already filled"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"updated_container": updatedID})
}

func getCore() (c *core) {
	c = coreInstance

	if c == nil {
		c = new(core)
		c.containerStorage = pg.Container{}
		coreInstance = c
	}

	return
}
