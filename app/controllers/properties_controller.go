package controllers

import (
	// "fmt"
	"github.com/benhawker/go-json-api/app/models"
	"github.com/benhawker/go-json-api/app/services"
	"github.com/revel/revel"
	"net/http"
)

type PropertiesController struct {
	*revel.Controller
	services.Database
}

func (c PropertiesController) Index() revel.Result {
	properties := make([]models.Property, 0)

	if err := c.Gorm.Find(&properties).Error; err != nil {
		panic(err)
	}
	return c.RenderJSON(properties)
}

func (c PropertiesController) Show(uuid string) revel.Result {
	property := models.Property{}
	c.Gorm.Where("uuid = ?", uuid).First(&property)

	if property.UUID == "" {
		c.Response.Status = http.StatusNotFound
		return c.RenderJSON("Property Not Found")
	}
	return c.RenderJSON(property)
}
