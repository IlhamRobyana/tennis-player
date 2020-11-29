package container

import (
	"github.com/ilhamrobyana/tennis-player/entity"
	pg "github.com/ilhamrobyana/tennis-player/pg_storage"
)

type core struct {
	containerStorage pg.Container
}

func (c *core) create(container entity.Container) (createdContainer entity.Container, e error) {
	client, e := pg.GetPGClient()
	defer client.Close()

	if e != nil {
		return entity.Container{}, e
	}
	c.containerStorage.Client = client
	createdContainer, e = c.containerStorage.Create(container)
	return
}

func (c *core) putBall(playerID uint64) (updatedID uint64, e error) {
	client, e := pg.GetPGClient()
	defer client.Close()

	if e != nil {
		return
	}
	c.containerStorage.Client = client
	updatedID, e = c.containerStorage.PutBall(playerID)
	return
}
