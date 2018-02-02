package controllers

import (
	"github.com/benhawker/go-json-api/app/models"
	"github.com/benhawker/go-json-api/app/services"
	"github.com/revel/revel"
)

type TimeSlotsController struct {
	*revel.Controller
	services.Database
}

func (c TimeSlotsController) Index() revel.Result {
	propertyUUID := c.Params.Route.Get("uuid")

	timeSlots := make([]models.TimeSlot, 0)

	if err := c.Gorm.Where("property_uuid = ?", propertyUUID).Find(&timeSlots).Error; err != nil {
		panic(err)
	}
	return c.RenderJSON(timeSlots)
}
