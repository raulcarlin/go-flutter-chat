// Run a web server for simple chat trough WebSockets.
package main

import (
	"backend/database"
	"backend/model"
	"backend/server"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.String("port", "8080", "Server port number")

func main() {
	flag.Parse()
	runDb()
	log.Println("Listening on port ", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", *port), server.Router()))
}

func runDb() {
	db, err := database.OpenDB()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	users, err := database.GetUsers(db)
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Printf("UID: %d - User: %s - Login: %s\n", user.UID, user.UserName, user.LastLogin.Format("02-Jan-2006"))
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	result, err := database.InsertUser(
		&model.User{UID: 1234567, UserName: "abcdef"}, tx)
	if err != nil {
		panic(err)
	}
	tx.Commit()

	fmt.Println(result.RowsAffected())
}
