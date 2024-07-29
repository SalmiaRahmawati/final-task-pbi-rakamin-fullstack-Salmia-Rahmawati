package main

import (
	"final-task-pbi-rakamin/database"
	"final-task-pbi-rakamin/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
