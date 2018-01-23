package controllers

import (
	"github.com/benhawker/go-json-api/app"
	"github.com/benhawker/go-json-api/app/models"
	// "github.com/benhawker/go-json-api/app/repository"
	"github.com/revel/revel"
)

type PropertiesController struct {
	*revel.Controller
}

func (c PropertiesController) Index() revel.Result {
	var properties []models.Property
	app.Gorm.Preload("Properties").Order("date asc")

	propertiesJSON := make([]models.PropertyJSON, len(properties))

	for i := range properties {
		propertiesJSON[i] = properties[i].ToJSON()
	}

	return c.RenderJSON(propertiesJSON)
}

// func (c PropertiesController) GetPropertyById(id string) revel.Result {

// 	user, err := repository.GetPropertyRepository().GetPropertyById(id)

// 	response := JsonResponse{}
// 	response.Success = err == nil
// 	response.Data = user
// 	if err != nil {
// 		response.Error = err.Error()
// 	}

// 	return c.RenderJson(response)
// }

// func (c PropertiesController) GetProperties() revel.Result {

// 	users := repository.GetPropertyRepository().GetProperties()

// 	response := JsonResponse{}
// 	response.Data = users

// 	return c.RenderJson(response)
// }

// func (c PropertiesController) SaveProperty(id, title string) revel.Result {

// 	user := &models.Property{
// 		Id:    id,
// 		Title: title,
// 	}

// 	err := repository.GetPropertyRepository().SaveProperty(user)

// 	response := JsonResponse{}
// 	response.Success = err == nil
// 	if err != nil {
// 		response.Error = err.Error()
// 	}
// 	return c.RenderJson(response)
// }
