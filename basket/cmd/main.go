package main

import (
	"log"
	"snippetBox-microservice/basket/routes"
)

func main() {

	r := routes.SetupRoutes()
	log.Println("listening on http://localhost:8000")
	r.Run(":8000")

}
