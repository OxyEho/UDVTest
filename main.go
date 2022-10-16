package main

import (
	"UDVTest/db"
	"UDVTest/handlers"
)

func main() {
	db.InitDb()
	repo := db.NewPostgresRepo()
	api := handlers.NewApi(&repo)
	api.Run()
}
