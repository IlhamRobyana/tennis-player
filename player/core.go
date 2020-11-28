package player

import (
	"errors"

	"github.com/ilhamrobyana/tennis-player/entity"
	"github.com/ilhamrobyana/tennis-player/helper"
	"github.com/ilhamrobyana/tennis-player/storage"
	"golang.org/x/crypto/bcrypt"
)

type core struct {
	playerStore    storage.PlayerStorage
	containerStore storage.ContainerStorage
}

func (c *core) signup(player entity.Player) (response entity.LoginResponse, e error) {
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

func (c *core) play(playerID uint64) (containers []entity.Container, e error) {
	containers, e = c.containerStore.GetFilledContainers(playerID)
	return
}
