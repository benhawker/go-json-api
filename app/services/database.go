package services

import (
	"fmt"
	"github.com/benhawker/go-json-api/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/revel/revel"
	"github.com/satori/go.uuid"
	"strconv"
	"time"
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
	if err != nil {

		panic(err)
	}

	db.SingularTable(true)
	return db
}

func InitDB() {
	db := ConnectToDatabase()
	db.DropTableIfExists(&models.Property{}, &models.TimeSlot{})
	db.CreateTable(&models.TimeSlot{})
	db.CreateTable(&models.Property{})
	db.AutoMigrate(&models.Property{}, &models.TimeSlot{})

	seedDB(db)
	db.Close()
}

func seedDB(db *gorm.DB) {
	for i := 0; i < 5; i++ {
		createProperty(db, i)
	}
}

func createProperty(db *gorm.DB, i int) {
	uuid, err := uuid.NewV4()

	if err != nil {
		panic(err)
	}

	propertyUUID := uuid.String()

	property := models.Property{UUID: propertyUUID, Title: fmt.Sprintf("Test Property %s", strconv.Itoa(i)), Description: "Another great test property", PostalCode: fmt.Sprintf("SW19 %sAB", strconv.Itoa(i)), PricePerMonth: 1000}
	db.NewRecord(property)
	db.Create(&property)
	db.NewRecord(property)

	// Create 3 timeSlots per property
	for j := 0; j < 3; j++ {
		createTimeSlot(db, i, propertyUUID)
	}
	return
}

func createTimeSlot(db *gorm.DB, i int, propertyUUID string) {
	exampleTime, err := time.Parse(time.RFC822, "21 Feb 18 10:00 UTC")
	if err != nil {
		panic(err)
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	tsUUID := uuid.String()

	ts := models.TimeSlot{UUID: tsUUID, DateTime: exampleTime, Available: true, PropertyUUID: propertyUUID}
	db.NewRecord(ts)
	db.Create(&ts)
	db.NewRecord(ts)
	return
}

func (c *Database) Open() revel.Result {
	c.Gorm = ConnectToDatabase()
	return nil
}

func (c *Database) Close() revel.Result {
	c.Gorm.Close()
	return nil
}
