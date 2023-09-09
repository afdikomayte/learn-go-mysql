package learn_go_mysql

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	//mulai transaction
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	//sql insert with prepare
	ctx := context.Background()

	sqlInsert := "INSERT INTO comments(email,comment) VALUES(?,?)"
	stmt, err := tx.PrepareContext(ctx, sqlInsert)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 10; i < 20; i++ {

		//data comment
		email := "trancation" + strconv.Itoa(i) + "@bcid.com"
		comment := "insert with transaction ke " + strconv.Itoa(i)

		//insert to db with prepare.execContext
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		//ambil last id insert
		lastId, _ := result.LastInsertId()
		fmt.Println("", lastId, comment, "oleh :", email)
	}

	// commit transaction
	tx.Rollback()
}
