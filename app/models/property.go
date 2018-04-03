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

type NotificationSubscription struct {
	Id           uint      `gorm:"primary_key" json:"id"`
	SubscriberId int       `json:"subscriber_id"`
	PublisherId  int       `json:"publisher_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// Plus add an index on requester and receiver.
}

type NotificationSubscriptionJSON struct {
	Type string `json:"type"`
	NotificationSubscription
}

func NewNotificationSubscriptionJSON(ns NotificationSubscription) NotificationSubscriptionJSON {
	nsJSON := NotificationSubscriptionJSON{"notification_subsription", ns}
	return nsJSON
}

type Block struct {
	Id          uint      `gorm:"primary_key" json:"id"`
	RequesterId int       `json:"requested_id"`
	BlockedId   int       `json:"blocked_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Plus add an index on requester and receiver.
}

type BlockJSON struct {
	Type string `json:"type"`
	Block
}

func NewBlockJSON(b Block) BlockJSON {
	blockJSON := BlockJSON{"block", b}
	return blockJSON
}
