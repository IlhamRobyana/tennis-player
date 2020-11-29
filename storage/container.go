package storage

import (
	"errors"

	"github.com/ilhamrobyana/tennis-player/entity"
	pg "github.com/ilhamrobyana/tennis-player/pg_storage"
)

type ContainerStorage interface {
	Create(container entity.Container) (entity.Container, error)
	PutBall(playerID uint64) (uint64, error)
	GetFilledContainer(playerID uint64) (entity.Container, error)
}

func GetContainerStorage(n int) (ContainerStorage, error) {
	switch n {
	case Postgre:
		return new(pg.Container), nil
	default:
		return nil, errors.New("not implemented yet")
	}
}
