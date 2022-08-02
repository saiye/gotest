package testcase

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	"time"
)

//https://gowebexamples.com/mysql-database/

func TestConnection(t *testing.T) {
	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "dev:123456!@#@(10.10.6.22:3306)/test1?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	err2 := db.Ping()
	if err2 != nil {
		log.Fatal(err2)
	}
	account := fmt.Sprintf("buffer %v", time.Now().Unix())
	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	// Inserts our data into the users table and returns with the result and a possible error.
	// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
	result, err3 := db.Exec(`INSERT INTO users (account,nickname, password, created_at,updated_at) VALUES (?, ?, ?,?,?)`, account, username, password, createdAt, createdAt)
	userID, err := result.LastInsertId()
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println(userID)
}
