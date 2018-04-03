package services

import (
	"fmt"
	"github.com/benhawker/go-json-api/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/revel/revel"
	// "github.com/lib/pq"
)

type Database struct {
	Gorm *gorm.DB
}

func ConnectToDatabase() *gorm.DB {
	connectionString := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		"localhost",
		"benh",
		"gojsonapitest",
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
	db.DropTableIfExists(&models.User{}, &models.Friendship{}, &models.NotificationSubscription{})
	db.CreateTable(&models.Friendship{})
	db.CreateTable(&models.User{})
	db.CreateTable(&models.NotificationSubscription{})
	db.CreateTable(&models.Block{})
	db.AutoMigrate(&models.User{}, &models.Friendship{}, &models.NotificationSubscription{})

	seedDB(db)
	db.Close()
}

func seedDB(db *gorm.DB) {
	for i := 0; i < 5; i++ {
		createUser(db, i)
		createFriendships(db, i)
		createNotificationSubs(db, i)
		createBlock(db, i)
	}
}

func createUser(db *gorm.DB, i int) {
	// user := models.User{Email: "test@email.com", BlockedList: bl}
	user := models.User{Email: "test@email.com"}
	db.NewRecord(user)
	db.Create(&user)
	db.NewRecord(user)

	// Create 3 timeSlots per property
	// for j := 0; j < 3; j++ {
	// 	createTimeSlot(db, i, propertyUUID)
	// }
	return
}

func createFriendships(db *gorm.DB, i int) {
	f := models.Friendship{RequesterId: i, ReceiverId: i + 1}
	db.NewRecord(f)
	db.Create(&f)
	db.NewRecord(f)
	return
}

func createNotificationSubs(db *gorm.DB, i int) {
	ns := models.NotificationSubscription{SubscriberId: i, PublisherId: i + 1}
	db.NewRecord(ns)
	db.Create(&ns)
	db.NewRecord(ns)
	return
}

func createBlock(db *gorm.DB, i int) {
	b := models.Block{RequesterId: i, BlockedId: i + 1}
	db.NewRecord(b)
	db.Create(&b)
	db.NewRecord(b)
	return
}

// func createTimeSlot(db *gorm.DB, i int, propertyUUID string) {

// 	ts := models.TimeSlot{UUID: tsUUID, DateTime: exampleTime, Available: true, PropertyUUID: propertyUUID}
// 	db.NewRecord(ts)
// 	db.Create(&ts)
// 	db.NewRecord(ts)
// 	return
// }

func (c *Database) Open() revel.Result {
	c.Gorm = ConnectToDatabase()
	return nil
}

func (c *Database) Close() revel.Result {
	c.Gorm.Close()
	return nil
}
