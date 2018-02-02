package controllers

import (
	"github.com/benhawker/go-json-api/app/models"
	"github.com/benhawker/go-json-api/app/services"
	"github.com/revel/revel"
	"net/http"
)

type TimeSlotsController struct {
	*revel.Controller
	services.Database
}

func (c TimeSlotsController) Index() revel.Result {
	propertyUUID := c.Params.Route.Get("uuid")

	property := models.Property{}
	c.Gorm.Where("uuid = ?", propertyUUID).First(&property)
	if property.UUID == "" {
		c.Response.Status = http.StatusNotFound
		resp := make(map[string]string)
		resp["mesage"] = "We don't have a property with the given UUID."
		return c.RenderJSON(resp)
	}

	timeSlots := make([]models.TimeSlot, 0)

	if err := c.Gorm.Where("property_uuid = ?", propertyUUID).Find(&timeSlots).Error; err != nil {
		panic(err)
	}
	return c.RenderJSON(timeSlots)
}

func (c TimeSlotsController) Create() revel.Result {
	ts := models.TimeSlot{}
	c.Params.BindJSON(&ts)

	c.Gorm.Create(&ts)
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(ts)
}
