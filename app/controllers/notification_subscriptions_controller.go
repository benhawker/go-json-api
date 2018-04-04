package controllers

import (
	"net/http"
	"fmt"

	"github.com/benhawker/go-json-api/app/models"
	"github.com/benhawker/go-json-api/app/services"
	"github.com/revel/revel"
)

type NotificationSubscriptionsController struct {
	*revel.Controller
	services.Database
}

func (c NotificationSubscriptionsController) Index() revel.Result {
	ns := make([]models.NotificationSubscription, 0)

	if err := c.Gorm.Where("subscriber_id = ? OR publisher_id = ? ", 1, 2).Find(&ns).Error; err != nil {
		c.RenderJSON(err)
	}

	json := make([]models.NotificationSubscriptionJSON, 0)
	for _, n := range ns {
		json = append(json, models.NewNotificationSubscriptionJSON(n))
	}
	return c.RenderJSON(json)
}


type NSRequestBody struct {
  Requestor string `json:"requestor"`
  Target string `json:"target"`
}

// Expected request body:
// {
//   "requestor": "lisa@example.com",
//   "target": "john@example.com"
// }

func (c NotificationSubscriptionsController) Create() revel.Result {
  var requestBody NSRequestBody
  c.Params.BindJSON(&requestBody)
  user := models.User{}

  // Check requester exists
  if err := c.Gorm.Where("email = ?", requestBody.Requestor).First(&user).Error; err != nil {
    c.Response.Status = http.StatusBadRequest
    return c.RenderJSON(fmt.Sprintf("We don't recognise the requesting user: %s", requestBody.Requestor))
  }

  // Check receiver exists
  if err := c.Gorm.Where("email = ?", requestBody.Target).First(&user).Error; err != nil {
    c.Response.Status = http.StatusBadRequest
    return c.RenderJSON(fmt.Sprintf("We don't recognise the target user: %s", requestBody.Target))
  }

  // Create the Notification Subscription
  // TODO

  // Render 200
  success := map[string]bool{"success": true}
  c.Response.Status = http.StatusOK
  return c.RenderJSON(success)
}
