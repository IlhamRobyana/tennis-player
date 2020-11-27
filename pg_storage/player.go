package pg_storage

import "github.com/ilhamrobyana/tennis-player/entity"

type Player struct{}

func (p *Player) Create(player entity.Player) (entity.Player, error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return entity.Player{}, e
	}
	e = client.Create(&player).Error
	return player, e
}

func (p *Player) GetByUsername(username string) (player entity.Player, e error) {
	client, e := GetPGClient()
	defer client.Close()
	if e != nil {
		return
	}
	e = client.Where("username = ?", username).First(&player).Error
	return
}
