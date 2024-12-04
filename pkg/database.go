package pkg

import (
	"database/sql"
	"fmt"
)

// It is caller's responsibility to close the connection once work is done.
func ConnectToDatabase(URL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", URL)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to database, err : %s", err.Error())
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error while pinging the database, err: %s", err.Error())
	}

	return db, nil
}
