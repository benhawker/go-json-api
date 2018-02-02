package models

import (
	"time"
)

type TimeSlot struct {
	UUID         string    `sql:"type:uuid;primary_key" json:"uuid"`
	DateTime     time.Time `json:"datetime"`
	Available    bool      `json:"available"`
	PropertyUUID string    `json:"property_uuid"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
