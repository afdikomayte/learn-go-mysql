package repository

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	connectionDB "github.com/afdikomayte/learn-go-mysql"
	"github.com/afdikomayte/learn-go-mysql/entity"
)

func TestCommentInsert(t *testing.T) {
	commentRepo := NewCommentRepository(connectionDB.GetConnection())

	//make context
	ctx := context.Background()

	//data comment entity
	comment := entity.Comment{
		Email:   "repo@bcid.com",
		Comment: "test insert from repository pattern",
	}
	result, err := commentRepo.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentUpdate(t *testing.T) {
	commentRepo := NewCommentRepository(connectionDB.GetConnection())

	//make context
	ctx := context.Background()

	//data untuk update
	id := 34
	comment := entity.Comment{
		Email:   "repoupdate@bcid.com",
		Comment: "Update from repo 2",
	}

	//script update repo
	result, err := commentRepo.Update(ctx, int32(id), comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepo := NewCommentRepository(connectionDB.GetConnection())

	//make parents context
	ctx := context.Background()
	// search id
	id := 33

	result, err := commentRepo.FindById(ctx, int32(id))
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestCommentFindAll(t *testing.T) {
	commentRepo := NewCommentRepository(connectionDB.GetConnection())

	//make parents context
	ctx := context.Background()

	comments, err := commentRepo.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		//cara hanya mengambil commentnya saja
		comment := comment.Comment
		fmt.Println(comment)
	}
}
