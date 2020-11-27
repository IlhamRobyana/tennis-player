package entity

// Container holds information of the container
type Container struct {
	ID       uint64 `json:"id" gorm:"primary_key;type:bigserial"`
	PlayerID uint64 `json:"player_id" gorm:"not null"`
	Capacity uint64 `json:"capacity" gorm:"not null"`
	Balls    uint64 `json:"balls" gorm:"not null"`
}

// ContainerCreateRequest holds information of what's needed to create a container
type ContainerCreateRequest struct {
	Capacity uint64 `json:"capacity" form:"capacity"`
}
