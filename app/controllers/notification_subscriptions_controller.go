package controllers

import (
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
		panic(err)
	}

	json := make([]models.NotificationSubscriptionJSON, 0)
	for _, n := range ns {
		json = append(json, models.NewNotificationSubscriptionJSON(n))
	}
	return c.RenderJSON(json)
}
