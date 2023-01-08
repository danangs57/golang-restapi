package conf

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Db is a package-level variable that holds the database connection
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/ecommerce")
	if err != nil {
		// handle the error
		panic(err)
	}

}
