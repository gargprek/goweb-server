package pkg

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // or the appropriate driver for your database
)

var DbConn *sql.DB

// GetDB returns the database connection
func GetDB() *sql.DB {
	return DbConn
}
