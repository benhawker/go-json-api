package models

import "time"

type Property struct {
	ID        uint `gorm:"primary_key"`
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PropertyJSON struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func (property *Property) ToJSON() PropertyJSON {
	return PropertyJSON{
		ID:    property.ID,
		Title: property.Title,
	}
}
