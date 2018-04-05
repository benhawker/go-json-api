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

	db.DropTableIfExists(
		&models.User{},
		&models.Friendship{},
		&models.NotificationSubscription{},
		&models.Block{},
		&models.Message{},
	)

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Friendship{})
	db.CreateTable(&models.NotificationSubscription{})
	db.CreateTable(&models.Block{})
	db.CreateTable(&models.Message{})

	db.AutoMigrate(
		&models.User{},
		&models.Friendship{},
		&models.NotificationSubscription{},
		&models.Block{},
		&models.Message{},
	)

	seedDB(db)
	db.Close()
}

func seedDB(db *gorm.DB) {
	for i := 1; i <= 6; i++ {
		createUser(db, i)
		createFriendships(db, i)
		createNotificationSubs(db, i)
		createBlock(db, i)
	}

	u := models.User{Email: "no_friends@email.com"}
	db.Create(&u)

	b := models.Block{RequesterId: 2, BlockedId: 1}
	db.Create(&b)

	b = models.Block{RequesterId: 3, BlockedId: 1}
	db.Create(&b)
}

func createUser(db *gorm.DB, i int) {
	user := models.User{Email: fmt.Sprintf("test_%v@email.com", i)}
	db.Create(&user)
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

func (c *Database) Open() revel.Result {
	c.Gorm = ConnectToDatabase()
	return nil
}

func (c *Database) Close() revel.Result {
	c.Gorm.Close()
	return nil
}
