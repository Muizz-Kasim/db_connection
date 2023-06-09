package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var db_conn *sql.DB
var err error

func Connect() {
	fmt.Println("Connecting to database...")

	// Mysql Config
	var hostName = "localhost"
	var port = "3306"
	var dbName = "recordings"

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 hostName + ":" + port,
		DBName:               dbName,
		AllowNativePasswords: true,
	}

	// Connect to database
	// db_conn_conn, err := sql.Open("mysql", "root:@tcp(localhost:3306)/progresif_db_conn")
	db_conn, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		// Handle error
		panic(err.Error())
	}

	// Test connection
	pingErr := db_conn.Ping()
	if pingErr != nil {
		// Handle error
		panic(err.Error())
	}
	fmt.Println("Connection to database: " + dbName + " established!")
	// defer db_conn.Close()
}
