package models

import (
	"github.com/revel/revel"
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

func (ts *TimeSlot) Validate(v *revel.Validation) {
	v.Required(ts.UUID)
	v.Required(ts.DateTime)
	v.Required(ts.Available)
	v.Required(ts.PropertyUUID)
	// v.Required(contains(allPropertyUUIDs(), ts.PropertyUUID)).MessageKey("We don't recognise the property you are attaching this time slot to.")
}

// func contains(slice []string, element string) bool {
// 	for _, a := range slice {
// 		if a == element {
// 			return true
// 		}
// 	}
// 	return false
// }

// func allPropertyUUIDs() []string {
// 	var propertyUUIDs []int64
// 	services.Database.Gorm.Find(&Property{}).Pluck("uuid", &propertyUUIDs)
// 	return propertyUUIDs
// }
