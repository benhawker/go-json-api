package models

import "time"
import "github.com/lib/pq"

type Property struct {
	UUID      string         `json:"uuid"`
	Title     string         `json:"title"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	TimeSlots pq.StringArray `gorm:"type:varchar(100)[]"` // Property has_many timeSlots
}

type PropertyJSON struct {
	UUID      string   `json:"uuid"`
	Title     string   `json:"title"`
	TimeSlots []string `json:"time_slots"`
}

func (property *Property) ToJSON() PropertyJSON {
	return PropertyJSON{
		UUID:      property.UUID,
		Title:     property.Title,
		TimeSlots: property.TimeSlots,
	}
}
