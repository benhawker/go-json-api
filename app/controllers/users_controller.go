package controllers

import (
	"net/http"

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
		c.Response.Status = http.StatusNotFound
		c.RenderJSON(err)
	}

	usersJson := make([]models.UserJSON, 0)
	for _, u := range users {
		usersJson = append(usersJson, models.NewUserJSON(u))
	}

	c.Response.Status = http.StatusOK
	return c.RenderJSON(usersJson)
}
