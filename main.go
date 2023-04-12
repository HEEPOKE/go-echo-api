package main

import (
	"fmt"
	"main/db"
	"main/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	db.ConnectDB()
	routes.Router()
}
