package database

import (
	"context"
	"fmt"
	"testing"
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

	for rows.Next(){
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}
