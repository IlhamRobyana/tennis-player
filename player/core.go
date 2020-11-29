package player

import (
	"errors"

	"github.com/ilhamrobyana/tennis-player/entity"
	"github.com/ilhamrobyana/tennis-player/helper"
	pg "github.com/ilhamrobyana/tennis-player/pg_storage"

	"golang.org/x/crypto/bcrypt"
)

type core struct {
	playerStore    pg.Player
	containerStore pg.Container
}

func (c *core) signup(player entity.Player) (response entity.LoginResponse, e error) {
	client, e := pg.GetPGClient()
	defer client.Close()
	if e != nil {
		return
	}
	c.playerStore.Client = client
	hashedPassword, e := bcrypt.GenerateFromPassword([]byte(player.Password), 10)
	if e != nil {
		return
	}
	player.Password = string(hashedPassword)
	createdPlayer, e := c.playerStore.Create(player)

	response.Token, e = helper.GenerateToken(createdPlayer)
	return
}

func (c *core) login(username, password string) (response entity.LoginResponse, e error) {
	client, e := pg.GetPGClient()
	defer client.Close()
	if e != nil {
		return
	}
	c.playerStore.Client = client
	player, e := c.playerStore.GetByUsername(username)
	if e != nil {
		return
	}
	e = bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(password))
	if e != nil {
		e = errors.New("Username or Password is wrong")
	}
	response.Token, e = helper.GenerateToken(player)
	return
}

func (c *core) play(playerID uint64) (container entity.Container, e error) {
	client, e := pg.GetPGClient()
	if e != nil {
		return
	}
	c.playerStore.Client = client
	c.containerStore.Client = client
	container, e = c.containerStore.GetFilledContainer(playerID)
	return
}
