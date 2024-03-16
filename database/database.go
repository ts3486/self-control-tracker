package database

import (
	"fmt"
	"os"

	"database/sql"
	"strconv"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var err error

func Connect(){
	host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    portStr := os.Getenv("DB_PORT")
	port, _ := strconv.Atoi(portStr)


	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, port)
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
}