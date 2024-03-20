package main

import (
	"MyGram/database"
	"MyGram/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
