package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() (*sql.DB, error) {
	connect := conString{}
	connect.prepareConnect()
	return sql.Open("postgres", connect.getConnectionString())
}

func SetDB(d *sql.DB) {
	db = d
}

func PingTest() error {
	return db.Ping()
}

func GetDB() *sql.DB {
	return db
}

const (
	disable = "disable"
	require = "require"
	ca      = "verify-ca"
	full    = "verify-full"
)

type sslMode string

type conString struct {
	name     string
	user     string
	password string
	host     string
	port     int
	ssl      sslMode
}

func (c *conString) prepareConnect() {
	c.fill_default()
	var err error
	c.host = os.Getenv("PGHOST")
	c.name = os.Getenv("PGDATABASE")
	c.password = os.Getenv("PGPASSWORD")
	c.port, err = strconv.Atoi(os.Getenv("PGPORT"))
	if err != nil {
		log.Println("Error, db port num cannot convert to type int.", err)
		os.Exit(1)
	}
	c.user = os.Getenv("PGUSER")
	c.ssl = disable
}

func (c *conString) fill_default() {
	if c.ssl == "" {
		c.ssl = disable
	}
	if c.port == 0 {
		c.port = 5432
	}
	if c.host == "" {
		c.host = "localhost"
	}
}

func (c conString) getConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=%s",
		c.name, c.user, c.password, c.host, c.port, c.ssl)
}

type Scanner interface {
	Scan(...interface{}) error
}

type Preparer interface {
	Prepare(query string) (*sql.Stmt, error)
}

/* func findQueringSqlIndex(sqlArray []string) (int, error) {
	scanFound := -1
	for i := 0; i < len(sqlArray) && scanFound >= 0; i++ {
		word := strings.ToUpper(sqlArray[i])
		if word == "SELECT" || word == "RETURNING" {

		}
	}
	if scanFound == -1 {
		return scanFound, fmt.Errorf("query key word not found")
	}
	return scanFound, nil
}

func sqlKeysList(sqlStatement string) ([]string, error) {
	var keys []string
	splitedSql := strings.Split(sqlStatement, " ")
	startIndex, err := findQueringSqlIndex(splitedSql)
	if err != nil {
		return nil, err
	}
	var endFound bool
	for i := startIndex + 1; i < len(splitedSql) && !endFound; i++ {
		keys = append(keys, splitedSql[i])
		if !strings.HasSuffix(splitedSql[i], ",") {
			endFound = true
		}
	}
	return keys, nil
}

func structDBTagList(s interface{}) error {
	t := reflect.TypeOf(s)
	strKind := t.Kind().String()
	if strKind != "Struct" {
		return fmt.Errorf("error require struct type instead got %s", strKind)
	}

	return nil
}
*/
