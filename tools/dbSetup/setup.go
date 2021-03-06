package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
	"github.com/ttacon/chalk"
)

var test bool
var drop bool

func init() {
	err := godotenv.Overload()
	if err != nil {
		log.Fatal("error load .env", err)
	}
	flag.BoolVar(&test, "test", false, "test mode")
	flag.BoolVar(&drop, "drop", false, "insert test data")
	flag.Parse()
}

func confirm() bool {
	var s string

	fmt.Printf("(y/N): ")
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}
	return false
}

func main() {
	scripts := []string{"uuid_ex.sql", "create_table.sql"}
	// if -test
	if test {
		scripts = append(scripts, "test_data.sql")
		// if -test -drop or -drop -test
		if drop {
			scripts = append(scripts, "drop_test.sql")
		}
		// if -drop
	} else if drop {
		fmt.Println(chalk.Red.Color("*********** Warning this will drop your entire system data. ************"))
		fmt.Println(chalk.Red.Color("you only want run this command when there is no real data in the system."))
		fmt.Println(chalk.Red.Color("---------------- do not run this command in production!! ----------------"))
		fmt.Println("are you sure want to drop your tables?")
		if !confirm() {
			fmt.Println("cancle")
			return
		}
		scripts = append(scripts, "drop_table.sql")
	}
	for _, script := range scripts {
		cmd := exec.Command("psql", "-f", "./dbSetup/"+script)
		err := cmd.Run()
		if err != nil {
			log.Fatal("error run script: " + script)
		}
	}
	log.Println("run success")
}
