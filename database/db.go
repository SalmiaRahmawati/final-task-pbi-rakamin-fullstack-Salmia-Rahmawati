package database

import (
	"final-task-pbi-rakamin/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "mysecret"
	port     = "5432"
	dbname   = "final_task"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	dsn := config
	var dbConnection *gorm.DB // untuk menyimpan hasil dari 'gorm .Open'
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connection to database :", err)
	}

	fmt.Println("sukses koneksi ke database")
	dbConnection.Debug().AutoMigrate(models.User{}, models.Photo{})

	// Assign variabel global db
	db = dbConnection
}

func GetDB() *gorm.DB {
	return db
}
