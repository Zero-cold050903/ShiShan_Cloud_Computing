package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/user"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	UID        string
	Username   string
	Password   string
	Permission string
}

func getUser() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}
func createDB() *sql.DB {
	usr := getUser()
	os.Mkdir(usr+"/.db", 0777)
	db, err := sql.Open("sqlite3", usr+"/.db/todo.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func createTable(db *sql.DB) {
	exec := "CREATE TABLE IF NOT EXISTS users (UID 			varchar(13) PRIMARY KEY," +
		"USERNAME 		varchar(20)," +
		"PASSWORD 		varchar(20))," +
		"PERMISSION 	varchar(20);"
	_, err := db.Exec(exec)
	if err != nil {
		log.Fatal(err)
	}
}
func InitDB() {
	db := createDB()
	defer db.Close()
	createTable(db)
	fmt.Println("Database created")
}
func InsertData(db *sql.DB, uid string, username string, password string, permission string) {
	exec := "INSERT INTO users (UID, USERNAME, PASSWORD, PERMISSION) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(exec, uid, username, password, permission)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data inserted")
}

var GET_SIZE = 100

func getData(db *sql.DB) []User {
	rows, err := db.Query("SELECT * FROM users")
	var users = make([]User, GET_SIZE)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		err := rows.Scan(&users[i].UID, &users[i].Username, &users[i].Password, &users[i].Permission)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return users
}
func inPassword() string {
	var password string
	fmt.Scanln(&password)
	return password
}
func inUsername() string {
	var username string
	fmt.Scanln(&username)
	return username
}
func matchKey(db *sql.DB) bool {
	for _, user := range getData(db) {
		if user.Username == inUsername() {
			return user.Password == inPassword()
		}
	}
	return false
}
