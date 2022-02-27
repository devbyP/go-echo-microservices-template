package main

import (
	"flag"
	"log"
	"os/exec"

	"github.com/joho/godotenv"
)

var test bool
var data bool

func init() {
	err := godotenv.Overload()
	if err != nil {
		log.Fatal("error load .env", err)
	}
	flag.BoolVar(&test, "test", false, "test mode")
	flag.BoolVar(&data, "data", false, "insert test data")
	flag.Parse()
}

func main() {
	scripts := [2]string{"uuid_ex.sql", "create_table.sql"}
	for _, script := range scripts {
		cmd := exec.Command("psql", "-f", "./dbSetup/"+script)
		err := cmd.Run()
		if err != nil {
			log.Fatal("error run script: " + script)
		}
	}
	log.Println("run success")
}
