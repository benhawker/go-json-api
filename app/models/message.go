package models

import (
	"github.com/lib/pq"
	"time"
)

type Message struct {
	Id         uint           `gorm:"primary_key" json:"id"`
	SenderId   int            `json:"requester_id"`
	Recipients pq.StringArray `gorm:"type:varchar(64)[]" json:"recipients"`
	Body       string         `json:"body"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	// Plus add an index on sender.
}
