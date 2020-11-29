package pg_storage

import (
	"math/rand"

	"github.com/ilhamrobyana/tennis-player/entity"
	"github.com/jinzhu/gorm"
)

type Container struct {
	Client *gorm.DB
}

func (c *Container) Create(container entity.Container) (entity.Container, error) {
	e := c.Client.Create(&container).Error
	return container, e
}

func (c *Container) PutBall(playerID uint64) (updatedID uint64, e error) {
	containerList := new([]entity.Container)

	e = c.Client.
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

	e = c.Client.
		Model(&updatingContainer).
		Where("id=?", updatedID).
		Updates(updatingContainer).
		Error
	return
}

func (c *Container) GetFilledContainer(playerID uint64) (container entity.Container, e error) {
	e = c.Client.
		Where("player_id=? AND balls = capacity", playerID).
		First(&container).
		Error
	if e != nil {
		return
	}
	e = c.Client.
		Model(&container).
		Where("id=?", container.ID).
		Update("balls", 0).
		Error
	return
}
