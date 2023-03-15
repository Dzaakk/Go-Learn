package database

import (
	"context"
	"fmt"
	"testing"
)

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
