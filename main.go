package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	log.Info(rows)
}

// Connect to a postgres db using the connectionstring
// `connectionString`.
// Performs a `SELECT 1` test query.
// Gives the single row test query result as result.
func connect(connectionSting string) (int, error) {
	db, err := sql.Open("postgres", connectionSting)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var resultValue int
	err = db.QueryRow("SELECT 1").Scan(&resultValue)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	fmt.Println(resultValue)
	return resultValue, nil
}
