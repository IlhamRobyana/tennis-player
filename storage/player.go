package storage

import (
	"errors"

	"github.com/ilhamrobyana/tennis-player/entity"
	pg "github.com/ilhamrobyana/tennis-player/pg_storage"
)

type PlayerStorage interface {
	Create(player entity.Player) (entity.Player, error)
	GetByUsername(username string) (entity.Player, error)
}

func GetPlayerStorage(n int) (PlayerStorage, error) {
	switch n {
	case Postgre:
		return new(pg.Player), nil
	default:
		return nil, errors.New("not implemented yet")
	}
}
