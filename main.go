package main

import (
	"sesi_12/database"
	"sesi_12/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8000")
}
