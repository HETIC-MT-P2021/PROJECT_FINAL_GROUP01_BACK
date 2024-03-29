package database

import (
	"database/sql"
	"fmt"
	"time"
	// Import for postgres
	_ "github.com/lib/pq"
)

var Database *sql.DB

// ConnectToDB Set up connection to the postgres DB
// Will panic on error
func ConnectToDB(host string, dbname string, user string, password string, port int64) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	tempDB, err := sql.Open("postgres", psqlInfo)

	// Open up our database connection.
	// set up a database on local machine using phpmyadmin.
	// The database is called gomvc

	if err != nil {
		fmt.Println("Database connection params error")
		panic(err)
	}

	err = tempDB.Ping()

	numberOfTest := 0

	for err != nil && numberOfTest < 5 {
		fmt.Println(err)
		fmt.Println("Connection to DB did not succeed, new try")

		time.Sleep(5 * time.Second)
		tempDB, err = sql.Open("postgres", psqlInfo)
		err = tempDB.Ping()

		numberOfTest++
	}

	if err != nil {
		fmt.Println("Database initialisation error")
		panic(err)
	}

	fmt.Println("Database successfully connected!")

	// defer the close till after the main function has finished
	// executing
	Database = tempDB
}

// CloseDbConnection will end dialogue with the DB
// Recommanded to use at program's end
func CloseDbConnection(db *sql.DB) {
	defer db.Close()
	fmt.Println("DB is closed")
}
