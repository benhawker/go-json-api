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
	f := make([]models.Friendship, 0)

	 c.Gorm.Where("requester_id = ? OR receiver_id = ? ", 1, 2).Find(&f)
	// 		panic(err)
	// }
	return c.RenderJSON(f)
}


