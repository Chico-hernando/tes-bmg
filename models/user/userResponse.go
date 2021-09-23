package user

import "time"

type UserResponse struct {
	Id                int       `json:"id"`
	Username          string    `json:"username"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	RefferalCode      string    `json:"refferalCode"`
	RefferalCodeInput string    `json:"refferalCodeInput"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
