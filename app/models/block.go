package models

import (
	"time"
)

type Block struct {
	Id          uint      `gorm:"primary_key" json:"id"`
	RequesterId int       `json:"requester_id"`
	BlockedId   int       `json:"blocked_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// Plus add an index on requester and blocked.
}

type BlockJSON struct {
	Type string `json:"type"`
	Block
}

func NewBlockJSON(b Block) BlockJSON {
	blockJSON := BlockJSON{"block", b}
	return blockJSON
}
