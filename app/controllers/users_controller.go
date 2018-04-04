package controllers

import (
	"github.com/benhawker/go-json-api/app/models"
	"github.com/benhawker/go-json-api/app/services"
	"github.com/revel/revel"
)

type UsersController struct {
	*revel.Controller
	services.Database
}

func (c UsersController) Index() revel.Result {
	users := make([]models.User, 0)

	if err := c.Gorm.Find(&users).Error; err != nil {
		c.RenderJSON(err)
	}

	json := make([]models.UserJSON, 0)
	for _, u := range users {
		json = append(json, models.NewUserJSON(u))
	}

	return c.RenderJSON(json)
}
