package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDBDefault() (*sql.DB, error) {
	connect := conString{}
	connect.prepareConnect()
	return ConnectDB(connect.String())
}

func ConnectDB(connectionString string) (*sql.DB, error) {
	return sql.Open("postgres", connectionString)
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
	Disable = iota
	Require
	Verify_ca
	Verify_full
)

type sslMode string

func setSSLMode(mode int) sslMode {
	modeString := [4]sslMode{"disable", "require", "verify_ca", "verify_full"}
	if mode > len(modeString) {
		return "disable"
	}
	return modeString[mode]
}

type conString struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     int
	ssl      sslMode
}

func NewConnectionString() *conString {
	c := new(conString)
	c.fill_default()
	return c
}

func (c *conString) SetSSL(mode int) {
	c.ssl = setSSLMode(mode)
}

func (c *conString) prepareConnect() {
	var err error
	c.Host = os.Getenv("PGHOST")
	c.Name = os.Getenv("PGDATABASE")
	c.Password = os.Getenv("PGPASSWORD")
	port := os.Getenv("PGPORT")
	if port == "" {
		port = "0"
	}
	c.Port, err = strconv.Atoi(port)
	if err != nil {
		log.Println("Error, db port cannot convert to type int.", err)
		os.Exit(1)
	}
	c.User = os.Getenv("PGUSER")
	c.fill_default()
}

func (c *conString) fill_default() {
	if c.ssl == "" {
		c.ssl = setSSLMode(Disable)
	}
	if c.Port == 0 {
		c.Port = 5432
	}
	if c.Host == "" {
		c.Host = "localhost"
	}
}

func (c conString) String() string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=%s",
		c.Name, c.User, c.Password, c.Host, c.Port, c.ssl)
}

type Scanner interface {
	Scan(...interface{}) error
}

type Preparer interface {
	Prepare(query string) (*sql.Stmt, error)
}

func ValidateUUID(id string) bool {
	uuidSectionNum := 5
	firstSectionLen := 8
	lastSectionLen := 12
	midSectionLen := 4
	sp := strings.Split(id, "-")
	if len(sp) != uuidSectionNum {
		return false
	} else if len(sp[0]) != firstSectionLen || len(sp[4]) != lastSectionLen {
		return false
	}
	for i := 1; i < 4; i++ {
		if len(sp[i]) != midSectionLen {
			return false
		}
	}
	return true
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
