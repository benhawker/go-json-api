package controllers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/benhawker/go-json-api/app/models"
	"github.com/benhawker/go-json-api/app/services"
	"github.com/revel/revel"
)

type MessagesController struct {
	*revel.Controller
	services.Database
}

// Expected request body:
// {
//   "sender":  "john@example.com",
//   "text": "Hello World! kate@example.com"
// }
type MessageRequestBody struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

// Expected response body:
// {
//   "success": true
//   "recipients":
//     [
//       "lisa@example.com",
//       "kate@example.com"
//     ]
// }
type MessageResponseBody struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}

func (c MessagesController) Create() revel.Result {
	var requestBody MessageRequestBody
	c.Params.BindJSON(&requestBody)

	sender := models.User{}
	if err := c.Gorm.Where("email = ?", requestBody.Sender).First(&sender).Error; err != nil {
		c.Response.Status = http.StatusNotFound
		return c.RenderJSON(fmt.Sprintf("We don't recognise the sender: %s", requestBody.Sender))
	}

	friendships := make([]models.Friendship, 0)
	if err := c.Gorm.Where("requester_id = ? OR receiver_id = ? ", sender.Id, sender.Id).Find(&friendships).Error; err != nil {
		panic(err)
	}

	subscriptions := make([]models.NotificationSubscription, 0)
	if err := c.Gorm.Where("subscriber_id = ?", sender.Id).Find(&subscriptions).Error; err != nil {
		panic(err)
	}

	userIds := []int{}

	// Include friends
	for _, f := range friendships {
		if sender.Id != uint(f.ReceiverId) {
			userIds = append(userIds, f.ReceiverId)
		}

		if sender.Id != uint(f.RequesterId) {
			userIds = append(userIds, f.RequesterId)
		}
	}

	// Include subscribers
	for _, s := range subscriptions {
		if sender.Id != uint(s.SubscriberId) {
			userIds = append(userIds, s.SubscriberId)
		}
	}

	// Include users mentioned in the message.
	r := regexp.MustCompile("[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+")
	matches := r.FindAllString(requestBody.Text, -1)

	fmt.Println(matches)

	user := models.User{}
	for _, email := range matches {
		if err := c.Gorm.Where("email = ?", email).First(&user).Error; err != nil {
			c.Response.Status = http.StatusNotFound
			return c.RenderJSON(fmt.Sprintf("We don't recognise an email in your text: %s", email))
		}

		userIds = append(userIds, int(user.Id))
	}

	// TODO: Remove Blocked Users
	// blocks := make([]models.Block{}, 0)
	// for _, uid := range userIds {
	// 	if err := c.Gorm.Select("requester_id").Where("requester_id = ? AND blocked_id", uid, sender.Id).Find(&blocks).Error; err != nil {
	// 		// Handle error
	// 		fmt.Println("User(s) was blocked - and removed.")
	// 	}
	// }

	// Get User records.
	users := make([]models.User, 0)
	if err := c.Gorm.Where("id in (?)", userIds).Find(&users).Error; err != nil {
		c.RenderJSON(err)
	}

	fmt.Println(userIds)

	// Create emails slice and append.
	emails := []string{}
	for _, u := range users {
		emails = append(emails, u.Email)
	}

	//Create the Message
	message := models.Message{
		SenderId:   int(sender.Id),
		Recipients: emails,
		Body:       requestBody.Text,
	}

	if err := c.Gorm.Create(&message).Error; err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(fmt.Sprintf("We could not send your message. Error Message: %s", err))
	}

	// Render 200
	c.Response.Status = http.StatusOK
	return c.RenderJSON(MessageResponseBody{true, emails})
}
