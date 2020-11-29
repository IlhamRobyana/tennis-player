package entity

// Player holds information of the player
type Player struct {
	ID       uint64 `json:"id" gorm:"primary_key;type:bigserial"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

// These following two structs are put on different structs
// despite having the same attributes because it serves different purposes
// and can hold different attributes if the circumstances were different

// SignupRequest holds information of request body on signup
type SignupRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// LoginRequest holds information of request body on login
type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// LoginResponse holds information of response body on success login
type LoginResponse struct {
	Token string `json:"token"`
}

// PlayResponse holds information of response body on success play
type PlayResponse struct {
	Message   string    `json:"message"`
	Container Container `json:"container"`
}
