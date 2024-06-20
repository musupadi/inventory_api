package Config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	// Data Source Name (DSN)
	user := "root"
	password := ""
	host := "127.0.0.1"
	port := "3306"
	database := "db_inventory"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Connected to the database!")
}

func Close() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Fatalf("Error closing the database: %v", err)
		}
		fmt.Println("Database connection closed.")
	} else {
		fmt.Println("Database connection was not initialized.")
	}
}
