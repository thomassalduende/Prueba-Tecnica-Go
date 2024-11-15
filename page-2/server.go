package main

import (
	"fmt"
	"log"
	"mutant-checker/database"
	"mutant-checker/routes"
	"net/http"
)

func main() {
	database.InitDB()
	defer database.CloseDB()

	routes.RegisterRoutes()
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
