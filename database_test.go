package learn_go_mysql_test

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabaseConnetion(t *testing.T) {
	db, err := sql.Open("mysql", "root:Password@db1@tcp(localhost:3306)/go_mysql")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
