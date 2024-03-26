package main

import (
	"fmt"
	"goPro/database"
	"goPro/routes"
)

func main(){
	err := database.ConnectDB()
	if err != nil {
			fmt.Println("Error connecting to database:", err)
			return
	}

	routes.SetupRoutes()
}