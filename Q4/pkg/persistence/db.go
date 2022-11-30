package persistence

import (
	common "byFood/Q4/pkg/common"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER = "burakakkaya"
	DB_NAME = "byFood"
)

// DB set up
func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", DB_USER, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	common.CheckError(err)

	return db
}
