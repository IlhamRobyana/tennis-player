package container

import (
	"github.com/ilhamrobyana/tennis-player/entity"
	"github.com/ilhamrobyana/tennis-player/storage"
)

type core struct {
	containerStorage storage.ContainerStorage
}

func (c *core) create(container entity.Container) (createdContainer entity.Container, e error) {
	createdContainer, e = c.containerStorage.Create(container)
	return
}

func (c *core) putBall(playerID uint64) (updatedID uint64, e error) {
	updatedID, e = c.containerStorage.PutBall(playerID)
	return
}
