package user

import "time"

type User struct {
	Id                int       `json:"id" gorm:"primaryKey"`
	Username          string    `json:"username" gorm:"unique"`
	Password          string    `json:"password"`
	Name              string    `json:"name"`
	Email             string    `json:"email" gorm:"unique"`
	RefferalCode      string    `json:"refferalCode"`
	RefferalCodeInput string    `json:"refferalCodeInput"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
