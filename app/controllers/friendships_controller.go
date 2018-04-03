package controllers

import (
	// "fmt"
	"github.com/benhawker/go-json-api/app/models"
	"github.com/benhawker/go-json-api/app/services"
	"github.com/revel/revel"
)

type FriendshipsController struct {
	*revel.Controller
	services.Database
}

func (c FriendshipsController) Index() revel.Result {
	f := make([]models.Friendship, 0)

	if err := c.Gorm.Where("requester_id = ? OR receiver_id = ? ", 1, 2).Find(&f).Error; err != nil {
		panic(err)
	}

	json := make([]models.FriendshipJSON, 0)
	for _, friendship := range f {
		json = append(json, models.NewFriendshipJSON(friendship))
	}

	return c.RenderJSON(json)
}

// func (c PropertiesController) Show() revel.Result {
// 	users := make([]models.User, 0)

// 	if err := c.Gorm.Find(&users).Error; err != nil {
// 		panic(err)
// 	}
// 	return c.RenderJSON(users)
// }

// func (c PropertiesController) Show(uuid string) revel.Result {
// 	property := models.Property{}
// 	c.Gorm.Where("uuid = ?", uuid).First(&property)

// 	if property.UUID == "" {
// 		c.Response.Status = http.StatusNotFound
// 		return c.RenderJSON("Property Not Found")
// 	}
// 	return c.RenderJSON(property)
// }
