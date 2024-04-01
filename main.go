package main

import (
	"fmt"
	"goProject/database.go"
	"goProject/routes"

	"github.com/gorilla/mux"
)

func main() {
	err := database.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	routes.SetupRoutes(r)
}
