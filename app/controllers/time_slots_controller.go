package controllers

import (
	"github.com/benhawker/go-json-api/app/models"
	"github.com/benhawker/go-json-api/app/services"
	"github.com/revel/revel"
	// "time"
)

type TimeSlotsController struct {
	*revel.Controller
	services.Database
}

func (c TimeSlotsController) Index() revel.Result {
	// exampleTime, err := time.Parse(time.RFC822, "21 Jan 18 10:00 UTC")

	// if err != nil {
	// 	panic(err)
	// }

	// ts := models.TimeSlot{UUID: "a1b2c3d4", TimeOfSlot: exampleTime, PropertyUUID: "a1b2c3d4"}
	// c.Gorm.NewRecord(ts) // => returns `true` as primary key is blank
	// c.Gorm.Create(&ts)
	// c.Gorm.NewRecord(ts) // => return `false` after `user` created

	timeSlots := make([]models.TimeSlot, 0)

	if err := c.Gorm.Find(&timeSlots).Error; err != nil {
		panic(err)
	}
	return c.RenderJSON(timeSlots)
}
