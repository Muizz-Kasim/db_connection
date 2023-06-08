package config

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql" // blank import
)

var helloWorld string

func Connect() {
	helloWorld = "Hello World"
	fmt.Println("Connecting to database...")

	// Mysql Config
	var hostName = "localhost"
	var port = "3306"

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 hostName + ":" + port,
		DBName:               "progresif_db",
		AllowNativePasswords: true,
	}

	// Connect to database
	// db_conn_conn, err := sql.Open("mysql", "root:@tcp(localhost:3306)/progresif_db_conn")
	db_conn, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		// Handle error
		panic(err.Error())
	}
	defer db_conn.Close()

	// Test connection
	pingErr := db_conn.Ping()
	if pingErr != nil {
		// Handle error
		panic(err.Error())
	}
	fmt.Println("Database connection established!")

}
