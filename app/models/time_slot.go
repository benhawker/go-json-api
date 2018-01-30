package models

import "time"

type TimeSlot struct {
	UUID       string `json:"uuid"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	TimeOfSlot time.Time

	// PropertyID is foreign key on timeSlot.
	Property     Property
	PropertyUUID string
}

type TimeSlotJSON struct {
	UUID         string    `json:"uuid"`
	PropertyUUID string    `json:"property_uuid"`
	TimeOfSlot   time.Time `json:"time_of_slot`
}

func (timeSlot *TimeSlot) ToJSON() TimeSlotJSON {
	return TimeSlotJSON{
		UUID:         timeSlot.UUID,
		PropertyUUID: timeSlot.PropertyUUID,
		TimeOfSlot:   timeSlot.TimeOfSlot,
	}
}
