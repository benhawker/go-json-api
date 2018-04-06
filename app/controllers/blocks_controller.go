package controllers

import (
	"fmt"
	"net/http"

	"github.com/benhawker/go-json-api/app/models"
	"github.com/benhawker/go-json-api/app/services"
	"github.com/revel/revel"
)

type BlocksController struct {
	*revel.Controller
	services.Database
}

func (c BlocksController) Index() revel.Result {
	b := make([]models.Block, 0)

	if err := c.Gorm.Find(&b).Error; err != nil {
		c.Response.Status = http.StatusNotFound
		return c.RenderJSON(err)
	}

	return c.RenderJSON(b)
}

// Expected request body:
// {
//   "requestor": "andy@example.com",
//   "target": "john@example.com"
// }
type BlockRequestBody struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

func (c BlocksController) Create() revel.Result {
	var requestBody BlockRequestBody
	c.Params.BindJSON(&requestBody)

	requestor := models.User{}
	target := models.User{}

	// // Check requester exists
	if err := c.Gorm.Where("email = ?", requestBody.Requestor).First(&requestor).Error; err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(fmt.Sprintf("We don't recognise the requesting user: %s", requestBody.Requestor))
	}

	// Check target exists
	if err := c.Gorm.Where("email = ?", requestBody.Target).First(&target).Error; err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(fmt.Sprintf("We don't recognise the target user: %s", requestBody.Target))
	}

	//Create the Block
	block := models.Block{RequesterId: int(requestor.Id), BlockedId: int(target.Id)}
	if err := c.Gorm.Create(&block).Error; err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(fmt.Sprintf("We could not save your Block Request. Error Message: %s", err))
	}

	// Render 200
	c.Response.Status = http.StatusOK
	return c.RenderJSON(ResponseBody{true})
}
