package database

import (
	"backend/model"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var (
	host   = "localhost"
	port   = 5432
	dbName = "postgres"
	user   = "admin"
	pass   = "P@ssw0rd"
)

func OpenDB() (db *sql.DB, err error) {
	connURL := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, pass, dbName)

	db, err = sql.Open("postgres", connURL)

	return
}

func GetUsers(db *sql.DB) (users []*model.User, err error) {
	users = make([]*model.User, 0)

	result, err := db.Query("SELECT uid, username, last_login FROM users")
	if err != nil {
		return
	}

	for result.Next() {
		user := &model.User{}
		result.Scan(&user.UID, &user.UserName, &user.LastLogin)
		users = append(users, user)
	}

	return
}

func InsertUser(u *model.User, tx *sql.Tx) (result sql.Result, err error) {
	result, err = tx.Exec("INSERT INTO users VALUES ($1, $2, $3)", u.UID, u.UserName, time.Now())

	return
}
