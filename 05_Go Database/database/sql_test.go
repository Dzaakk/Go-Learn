package database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

//eksekusi bukan query
func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	InputScript := "INSERT INTO customer(id, name) VALUES('rani', 'Rani')"
	_, err := db.ExecContext(ctx, InputScript)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert New Customer")
}

//query context untuk eksekusi
func TestQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

func TestQuerySQLComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("-------------")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createAt)
	}
}
