package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/afdikomayte/learn-go-mysql/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email,comment) VAlUES (?,?)"
	result, err := repo.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil

}

func (repo *commentRepositoryImpl) Update(ctx context.Context, id int32, comment entity.Comment) (entity.Comment, error) {
	script := "UPDATE comments SET email = ?, comment = ? WHERE id = ?"

	_, err := repo.DB.ExecContext(ctx, script, comment.Email, comment.Comment, id)
	if err != nil {
		return comment, err
	}

	comment.Id = id

	return comment, nil

}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id,email,comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, script, id)
	//make variabel comment value entity comment
	comment := entity.Comment{}
	if err != nil {
		return comment, nil
	}
	defer rows.Close()

	// cek rows.Next() jika true maka data ada
	if rows.Next() {

		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}
func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id,email,comment FROM comments"

	rows, err := repo.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment

	//memasukan data dari Database ke var comments
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil

}
