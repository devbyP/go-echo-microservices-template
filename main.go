package main

import (
	"log"
	"os"

	"github.com/devbyP/untitled/pkg/models"
	"github.com/devbyP/untitled/server"
	"github.com/joho/godotenv"
)

func loadEnv() {
	var err error
	if os.Getenv("GO_DEV") == "development" {
		err = godotenv.Overload()
	} else {
		err = godotenv.Load()
	}
	if err != nil {
		log.Fatal("error load dotenv file")
	}
}

func main() {
	loadEnv()
	port := os.Getenv("PORT")

	db, err := models.ConnectDBDefault()
	if err != nil {
		log.Fatal("cannot connect to database")
	}
	defer db.Close()

	models.SetDB(db)

	server.SetPort(port)
	server.StartServer()
}
