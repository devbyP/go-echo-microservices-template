package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID        string    `db:"id" json:"id"`
	FirstName string    `db:"first_name" json:"firstName"`
	LastName  string    `db:"last_name" json:"lastName"`
	Username  string    `db:"username" json:"username"`
	JoinDate  time.Time `db:"join_date" json:"joinDate"`
}

func NewUser(fn, ln, un string) *User {
	return &User{
		ID:        "",
		FirstName: fn,
		LastName:  ln,
		Username:  un,
		JoinDate:  time.Now(),
	}
}

func prepareUserSelectStmt(option string) (*sql.Stmt, error) {
	sel := "SELECT id, first_name, last_name, username, join_date FROM users"
	if strings.HasPrefix(strings.ToUpper(option), "WHERE") {
		return db.Prepare(sel + " " + option)
	}
	return db.Prepare(sel)
}

func prepareUserInsertStmt(tx Preparer, returning string) (*sql.Stmt, error) {
	ins := `INSERT INTO users (
		id, 
		first_name, 
		last_name, 
		username, 
		hash, 
		join_date
	) VALUES (uuid_generate_v4(), $1, $2, $3, $4, now())`
	if strings.HasPrefix(strings.ToUpper(returning), "RETURNING") {
		return tx.Prepare(ins + " " + returning)
	}
	return tx.Prepare(ins)
}

func (u *User) scanUser(row Scanner) error {
	return row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Username, &u.JoinDate)
}

func (u *User) fetchRow(option string, args ...interface{}) error {
	stmt, err := prepareUserSelectStmt(option)
	if err != nil {
		return err
	}
	defer stmt.Close()
	var row *sql.Row
	if len(args) == 0 {
		row = stmt.QueryRow()
	} else {
		row = stmt.QueryRow(args...)
	}
	err = u.scanUser(row)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil
		}
		return err
	}
	return nil
}

func useUserFetchRow(option string, args ...interface{}) (*User, error) {
	user := new(User)
	err := user.fetchRow(option, args...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func FetchUser(id string) (*User, error) {
	return useUserFetchRow("WHERE id = $1", id)
}

func FetchUserByUsername(uname string) (*User, error) {
	return useUserFetchRow("WHERE username = $1", uname)
}

func fetchMultiple(option string, args ...interface{}) ([]*User, error) {
	var err error
	users := make([]*User, 0)
	stmt, err := prepareUserSelectStmt(option)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var rows *sql.Rows
	if len(args) == 0 {
		rows, err = stmt.Query()
	} else {
		rows, err = stmt.Query(args...)
	}
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := new(User)
		err = user.scanUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func FetchAllUser() ([]*User, error) {
	return fetchMultiple("")
}

func (u *User) SignUp(hash string) error {
	instmt, err := prepareUserInsertStmt(db, "RETURNING id")
	if err != nil {
		return err
	}
	defer instmt.Close()
	row := instmt.QueryRow(u.FirstName, u.LastName, u.Username, hash)
	err = row.Scan(&u.ID)
	return err
}

func GetUserHash(id string) (string, error) {
	row := db.QueryRow("SELECT hash FROM users WHERE id = $1")
	var hash string
	if err := row.Scan(&hash); err != nil {
		return "", fmt.Errorf("error, cannot get hash from user that does not exist")
	}
	return hash, nil
}
