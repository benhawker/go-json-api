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

	if err := c.Gorm.Find(&f).Error; err != nil {
		// if err := c.Gorm.Where("requester_id = ? OR receiver_id = ? ", 999, 888).Find(&f).Error; err != nil {
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

// {
//   friends:
//     [
//       'andy@example.com',
//       'john@example.com'
//     ]
// }

type RequestBody struct {
	Friends []string `json:"friends"`
}

func (c FriendshipsController) Create() revel.Result {
	var jsonData RequestBody
	c.Params.BindJSON(&jsonData)

	users := make([]models.User, 0)

	// Check requester exists
	// req := models.User{Email: jsonData.Friends[0]}
	if err := c.Gorm.Where("email = ?", jsonData.Friends[0]).Find(&users).Error; err != nil {
		return c.RenderJSON("We don't recognise the requesting user.")
	}

	if len(users) == 0 {
		return c.RenderJSON("No users")
	}

	// Check receiver exists
	// rec := models.User{Email: jsonData.Friends[1]}
	// if err := c.Gorm.Find(&rec).Error; err != nil {
	// 	return c.RenderJSON("We don't recognise the receiving user.")
	// }

	success := map[string]bool{"success": true}
	return c.RenderJSON(success)

	// if err := c.Gorm.Find(&users).Error; err != nil {
	// 	panic(err)
	// }

	// f := models.Friendship{RequesterId: 999, ReceiverId: 888}

	// if err := c.Gorm.Create(&f).Error; err != nil {
	// 	return c.RenderJSON("custom error msg.")
	// }

	// // json := make([]models.FriendshipJSON, 0)
	// // for _, friendship := range f {
	// // 	json = append(json, models.NewFriendshipJSON(friendship))
	// // }

	// return c.RenderJSON(f)
}
