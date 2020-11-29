package pg_storage

import (
	"math/rand"

	"github.com/ilhamrobyana/tennis-player/entity"
)

type Container struct{}

func (c *Container) Create(container entity.Container) (entity.Container, error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return entity.Container{}, e
	}

	e = client.Create(&container).Error
	return container, e
}

func (c *Container) PutBall(playerID uint64) (updatedID uint64, e error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return
	}

	containerList := new([]entity.Container)

	e = client.
		Where("player_id=? AND balls < capacity", playerID).
		Find(&containerList).
		Error
	if e != nil || len(*containerList) == 0 {
		return
	}

	element := uint64(rand.Intn(len(*containerList)))
	updatingContainer := (*containerList)[element]
	updatingContainer.Balls++
	updatedID = updatingContainer.ID

	e = client.
		Model(&updatingContainer).
		Where("id=?", updatedID).
		Updates(updatingContainer).
		Error
	return
}

func (c *Container) GetFilledContainer(playerID uint64) (container entity.Container, e error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return
	}
	e = client.
		Where("player_id=? AND balls = capacity", playerID).
		First(&container).
		Error
	if e != nil {
		return
	}
	e = client.
		Model(&container).
		Where("id=?", container.ID).
		Update("balls", 0).
		Error
	return
}
