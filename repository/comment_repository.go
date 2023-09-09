package repository

import (
	"context"

	entity "github.com/afdikomayte/learn-go-mysql/entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	Update(ctx context.Context, id int32, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
	// Deleteee(ctx context.Context,id int32)
}
