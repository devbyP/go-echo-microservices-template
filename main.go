package main

import (
	"log"
	"os"

	"github.com/devbyP/untitled/models"
	"github.com/devbyP/untitled/server"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Overload()
	if err != nil {
		log.Fatal("error load dotenv file")
	}
}

func main() {
	port := os.Getenv("PORT")

	db, err := models.ConnectDB()
	if err != nil {
		log.Fatal("cannot connect to database")
	}
	defer db.Close()

	models.SetDB(db)

	server.SetPort(port)
	log.Println("server start at port: " + port)
	server.StartServer()
}