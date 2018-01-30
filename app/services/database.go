package services

import (
	"fmt"
	"github.com/benhawker/go-json-api/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/revel/revel"
)

type Database struct {
	Gorm *gorm.DB
}

func ConnectToDatabase() *gorm.DB {
	connectionString := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		"localhost",
		"benh",
		"gojsonapi",
		"disable",
		"password")

	db, err := gorm.Open("postgres", connectionString)
	fmt.Println("Big issue connecting.")
	if err != nil {

		panic(err)
	}

	db.SingularTable(true)

	return db
}

func InitDB() {
	db := ConnectToDatabase()
	db.AutoMigrate(&models.Property{}, &models.TimeSlot{})
	db.Close()
}

func (c *Database) Open() revel.Result {
	c.Gorm = ConnectToDatabase()
	return nil
}

func (c *Database) Close() revel.Result {
	c.Gorm.Close()
	return nil
}
