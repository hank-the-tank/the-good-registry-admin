package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "hank"
	password = "password"
	dbname   = "the_good_registry"
)

func db() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

// func dbMigration() {
// 	func main() {
// 		db, err := sql.Open("postgres", "postgres://localhost:5432/database?sslmode=enable")
// 		driver, err := postgres.WithInstance(db, &postgres.Config{})
// 		m, err := migrate.NewWithDatabaseInstance(
// 			"file:///migrations",
// 			"postgres", driver)
// 		m.Steps(2)
// 	}
// }
