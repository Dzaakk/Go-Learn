package repository

import (
	"context"
	"entity"
	"fmt"
	DBcon "go-database"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	CommentRepository := NewCommentRepository(DBcon.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repositoryc",
	}
	result, err := CommentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}
func TestFindById(t *testing.T) {
	CommentRepository := NewCommentRepository(DBcon.GetConnection())
	comment, err := CommentRepository.FindbyId(context.Background(), 37)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}
func TestFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(DBcon.GetConnection())
	comments, err := CommentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {

		fmt.Println(comment)
	}

}
