package repository

// import (
// 	"errors"
// 	"github.com/benhawker/go-json-api/app/models"
// )

// type PropertyRepository interface {
// 	GetProperties() []*models.Property
// 	GetPropertyById(id string) (*models.Property, error)
// 	SaveProperty(user *models.Property) error
// }

// type DBPropertyRepository struct {
// 	properties []*models.Property
// }

// func New() *DBPropertyRepository {
// 	return &DBPropertyRepository{
// 		properties: []*models.Property{
// 			&models.Property{1, "Test Property 1"},
// 			&models.Property{2, "Test Property 2"},
// 			&models.Property{3, "Test Property 3"},
// 			&models.Property{4, "Test Property 4"},
// 		},
// 	}
// }

// func (r *DBPropertyRepository) GetProperties() []*models.Property {
// 	return r.properties
// }

// func (r *DBPropertyRepository) GetPropertyById(id string) (*models.Property, error) {
// 	for _, property := range r.properties {
// 		if property.Id == id {
// 			return property, nil
// 		}
// 	}
// 	return nil, errors.New("Property not found.")
// }

// func (r *DBPropertyRepository) SaveProperty(property *models.Property) error {
// 	r.properties = append(r.properties, property)
// 	return nil
// }

// var propertyRepository *DBPropertyRepository

// func GetPropertyRepository() (r PropertyRepository) {
// 	if propertyRepository == nil {
// 		propertyRepository = New()
// 	}
// 	return propertyRepository
// }
