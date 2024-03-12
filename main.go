package main

import (
	"assignment2/database"
	"assignment2/routes"

	"fmt"
)

func main() {
	fmt.Println("Belajar Golang")
	database.StartDB()
	r := routes.SetupRoutes()
	r.Run(":8000")

}
