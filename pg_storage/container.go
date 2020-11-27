package pg_storage

import "github.com/ilhamrobyana/tennis-player/entity"

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
