package pg_storage

import (
	"github.com/ilhamrobyana/tennis-player/entity"
	"github.com/jinzhu/gorm"
)

type Player struct {
	Client *gorm.DB
}

func (p *Player) Create(player entity.Player) (entity.Player, error) {
	e := p.Client.Create(&player).Error
	return player, e
}

func (p *Player) GetByUsername(username string) (player entity.Player, e error) {
	e = p.Client.Where("username = ?", username).First(&player).Error
	return
}
