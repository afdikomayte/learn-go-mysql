package learn_go_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlScript := "INSERT INTO customer(id,nama) VALUES ('C3','Annta Rispo');"
	_, err := db.ExecContext(ctx, sqlScript)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data to database")
}

func TestQueryContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlSelect := "SELECT * FROM customer;"
	rows, err := db.QueryContext(ctx, sqlSelect)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, nama string
		err := rows.Scan(&id, &nama)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id :", id)
		fmt.Println("Nama :", nama)
	}
	defer rows.Close()

}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlSelect := "SELECT id,name,email,balance,rating,birth_date,merried,created_at FROM customer_detail"
	rows, err := db.QueryContext(ctx, sqlSelect)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString //if feild in database can be NULL
		var balance int32
		var rating float64 //double
		var merried bool
		var birthDate sql.NullTime //if feild in database can be NULL
		var createdAt time.Time

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &merried, &createdAt)
		if err != nil {
			panic(err)
		}
		//cek email & birth date because them NULLABLE
		if email.Valid || birthDate.Valid {
			fmt.Println("=======================")
			fmt.Println("Id :", id, "Name :", name, "Email : ", email.String, "Balance :", balance, "Rating :", rating, "Merried :", merried, "Birth Date :", birthDate.Time, "Created At :", createdAt)
		} else {
			fmt.Println("=======================")
			fmt.Println("Id :", id, "Name :", name, "Email : -", "Balance :", balance, "Rating :", rating, "Merried :", merried, "Birth Date : -", "Created At :", createdAt)
		}

	}
	defer rows.Close()
}

func TestQueryWithParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	//username and password for select
	username := "afdiko"
	password := "password"

	ctx := context.Background()

	sqlParamScript := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, sqlParamScript, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var userProfileName string
		err := rows.Scan(&userProfileName)
		if err != nil {
			panic(err)
		}

		fmt.Println("login berhasil ! welcome ", userProfileName)
	} else {
		fmt.Println("login gagal ! username atau password salah ")
	}
	rows.Close()
}

// melakukan operasi sql dengan Prepare Context dan melakukan insert banyak data agar tidak banyak koneksi
// pooling db yang digunakan karena prepare akan mengingat koneksi dengan sql yang dilakaukan hanya parameter
// yang berbeda
func TestPrepareContextAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlInsert := "INSERT INTO comments(email,comment) VALUES(?,?)"
	stmt, err := db.PrepareContext(ctx, sqlInsert)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		//data comment
		email := "afdiko" + strconv.Itoa(i) + "@bcid.com"
		comment := "koment ke " + strconv.Itoa(i)

		//insert param to stmt prepare
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("last insert Id : ", insertId)
	}
}
