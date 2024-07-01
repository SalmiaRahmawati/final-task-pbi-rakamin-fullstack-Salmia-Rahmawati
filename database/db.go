package database

import (
	"final-task-pbi-rakamin/app"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	// err      error
)

func StartDB() {
	host := "localhost"
	user := "postgres"
	password := "mysecret"
	port := "5432"
	dbname := "final_task"

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	dsn := config
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connection to database :", err)
	}

	if err := db.AutoMigrate(&app.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	DB = db
}

// var (
// 	host     = "localhost"
// 	user     = "postgres"
// 	password = "mysecret"
// 	port     = "5432"
// 	dbname   = "final_task"
// 	db       *gorm.DB
// )

// func StartDB() {
// 	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
// 	dsn := config
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal("error connection to database :", err)
// 	}

// 	fmt.Println("sukses koneksi ke database")
// 	db.Debug().AutoMigrate(app.User{}, app.Photo{})
// }

// func GetDB() *gorm.DB {
// 	return db
// }

// 	fmt.Println("sukses koneksi ke database")
// 	db.Debug().AutoMigrate(app.User{}, app.Photo{})
// }

// 	fmt.Println("sukses koneksi ke database")
// 	db.Debug().AutoMigrate(app.User{}, app.Photo{})
// }

// DB = db

// func DB() *gorm.DB {
// 	host := "localhost"
// 	user := "postgres"
// 	password := "mysecret"
// 	port := "5432"
// 	dbname := "final_task"

// 	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
// 	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
// 	if err != nil {
// 		panic("error connection to database")
// 	}
// 	return db
// }
