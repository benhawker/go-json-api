package controllers

import (
	"github.com/benhawker/go-json-api/app/models"
	"github.com/benhawker/go-json-api/app/services"
	"github.com/revel/revel"
)

type PropertiesController struct {
	*revel.Controller
	services.Database
}

func (c PropertiesController) Index() revel.Result {
	// property := models.Property{UUID: "a1b2c3d4", Title: "Test", TimeSlots: []string{"1", "2", "3", "4"}}
	// c.Gorm.NewRecord(property) // => returns `true` as primary key is blank
	// c.Gorm.Create(&property)
	// c.Gorm.NewRecord(property) // => return `false` after `user` created

	properties := make([]models.Property, 0)

	if err := c.Gorm.Find(&properties).Error; err != nil {
		panic("Houston")
		panic(err)
	}
	return c.RenderJSON(properties)
}
