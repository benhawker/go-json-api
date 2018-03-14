package models

import (
	"time"
	"github.com/lib/pq"
)

type User struct {
	Id uint `gorm:"primary_key"`
	Email       string `json:"email"`

	BlockedList  pq.StringArray `gorm:"type:varchar(64)[]" json:"block_list"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type Friendship struct {
	Id 					uint `gorm:"primary_key"`
	RequesterId         int `json:"requester_id"`
	ReceiverId          int `json:"receiver_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	// Plus add an index on requester and receiver.
}

type NotificationSubscription struct {
	Id 					uint `gorm:"primary_key"`
	SubscriberId         int `json:"subscriber_id"`
	PublisherId          int `json:"publisher_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	// Plus add an index on requester and receiver.
}

type Block struct {
	Id 					uint `gorm:"primary_key"`
	RequesterId         int `json:"requested_id"`
	BlockedId          int `json:"blocked_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	// Plus add an index on requester and receiver.
}