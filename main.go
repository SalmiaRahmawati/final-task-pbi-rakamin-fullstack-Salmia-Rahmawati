package main

import (
	"final-task-pbi-rakamin/database"
	"final-task-pbi-rakamin/router"
)

func main() {
	database.StartDB()
	// r := gin.Default()
	// // Mengatur IP proxy yang dipercaya secara eksplisit
	// err := r.SetTrustedProxies([]string{"192.168.1.2", "192.168.1.3"}) // Ganti dengan IP proxy yang dipercaya
	// if err != nil {
	// 	log.Fatalf("failed to set trusted proxies: %v", err)
	// }

	r := router.StartApp()
	r.Run(":8080")
}
