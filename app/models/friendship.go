package models

import (
	"time"
)

type Friendship struct {
	Id          uint      `gorm:"primary_key" json:"id"`
	RequesterId int       `json:"requester_id"`
	ReceiverId  int       `json:"receiver_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// Plus add an index on requester and receiver.
}

type FriendshipJSON struct {
	Type string `json:"type"`
	Friendship
}

func NewFriendshipJSON(f Friendship) FriendshipJSON {
	friendshipJSON := FriendshipJSON{"friendship", f}
	return friendshipJSON
}
