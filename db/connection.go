package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {
	dsn := "root:mysql@2025@tcp(127.0.0.1:3306)/task_manager"

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Failed to connect to database:  %v\n", err)
	}
	if err := DB.Ping(); err != nil {
		fmt.Printf("Failed to ping DB: %v\n", err)
	}

	fmt.Printf("Connected to Mysql\n")
	return nil
}
