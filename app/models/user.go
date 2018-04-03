package models

import (
	"github.com/lib/pq"
	"time"
)

type User struct {
	Id    uint   `gorm:"primary_key" json:"id"`
	Email string `json:"email"`

	BlockedList pq.StringArray `gorm:"type:varchar(64)[]" json:"block_list"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type UserJSON struct {
	Type string `json:"type"`
	User
}

func NewUserJSON(u User) UserJSON {
	userJSON := UserJSON{"user", u}
	return userJSON
}
