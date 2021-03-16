package main

import (
	"log"
	"snippetBox-microservice/basket/connection"
	"snippetBox-microservice/basket/routes"
)

func main() {

	db := connection.SetupDB()
	r := routes.SetupRoutes(db)
	log.Println("listening on http://localhost:8000")
	r.Run(":8000")

}
