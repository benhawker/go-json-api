package controllers

import (
	"net/http"

	"fmt"
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
		c.Response.Status = http.StatusNotFound
		return c.RenderJSON(err)
	}

	json := make([]models.FriendshipJSON, 0)
	for _, friendship := range f {
		json = append(json, models.NewFriendshipJSON(friendship))
	}

	return c.RenderJSON(json)
}

// GET http://localhost:9000/friendships/test@email.com
//
// Expected response body:
//
// {
//   "success": true,
//   "friends" :
//     [
//       'john@example.com'
//     ],
//   "count" : 1
// }
//
type FriendshipResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

func (c FriendshipsController) Show(email string) revel.Result {
	user := models.User{}

	if err := c.Gorm.Where("email = ?", email).First(&user).Error; err != nil {
		c.Response.Status = http.StatusNotFound
		return c.RenderJSON(fmt.Sprintf("We have no user by this email: %s", email))
	}

	friendships := make([]models.Friendship, 0)
	if err := c.Gorm.Where("requester_id = ? OR receiver_id = ? ", user.Id, user.Id).Find(&friendships).Error; err != nil {
		return c.RenderJSON(err)
	}

	userIds := []int{}
	for _, f := range friendships {
		if user.Id != uint(f.ReceiverId) {
			userIds = append(userIds, f.ReceiverId)
		}

		if user.Id != uint(f.RequesterId) {
			userIds = append(userIds, f.RequesterId)
		}
	}

	users := make([]models.User, 0)
	if err := c.Gorm.Where("id in (?)", userIds).Find(&users).Error; err != nil {
		c.Response.Status = http.StatusBadRequest
		c.RenderJSON(err)
	}

	emails := []string{}
	for _, u := range users {
		emails = append(emails, u.Email)
	}

	// Render 200
	c.Response.Status = http.StatusOK
	return c.RenderJSON(FriendshipResponse{Success: true, Friends: emails, Count: len(users)})
}

type RequestBody struct {
	Friends []string `json:"friends"`
}

// Expected request body:
// {
//   friends:
//     [
//       'andy@example.com',
//       'john@example.com'
//     ]
// }
func (c FriendshipsController) Create() revel.Result {
	var requestBody RequestBody
	c.Params.BindJSON(&requestBody)

	requestor := models.User{}
	receiver := models.User{}

	// Check requester exists
	if err := c.Gorm.Where("email = ?", requestBody.Friends[0]).First(&requestor).Error; err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(fmt.Sprintf("We don't recognise the requesting user: %s", requestBody.Friends[0]))
	}

	// Check receiver exists
	if err := c.Gorm.Where("email = ?", requestBody.Friends[1]).First(&receiver).Error; err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(fmt.Sprintf("We don't recognise the receiving user: %s", requestBody.Friends[1]))
	}

	// If the user is blocked then the friends connection cannot be made.
	block := models.Block{}
	c.Gorm.Where("requester_id = ? AND blocked_id = ?", requestor.Id, receiver.Id).First(&block)
	if block == (models.Block{}) {
		c.Response.Status = http.StatusForbidden
		return c.RenderJSON(fmt.Sprintf("The receiver of this friend request has blocked the requestor"))
	}

	//Create the Friendship
	friendship := models.Friendship{RequesterId: int(requestor.Id), ReceiverId: int(receiver.Id)}
	if err := c.Gorm.Create(&friendship).Error; err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(fmt.Sprintf("We could not save your Friend Request. Error Message: %s", err))
	}

	// Render 200
	c.Response.Status = http.StatusOK
	return c.RenderJSON(ResponseBody{true})
}
