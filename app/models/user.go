package models

import (
	"time"
)

type User struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserJSON struct {
	Type string `json:"type"`
	User
}

func NewUserJSON(u User) UserJSON {
	userJSON := UserJSON{"user", u}
	return userJSON
}
