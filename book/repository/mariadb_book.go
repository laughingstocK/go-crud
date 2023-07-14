package repository

import (
	"context"
	"database/sql"

	"github.com/laughingstocK/go-crud/book"
	"github.com/laughingstocK/go-crud/models"
)

type mariadbBookRepo struct {
	Conn *sql.Conn
}

func NewMariadbBookRepo(Conn *sql.Conn) book.Repository {
	return &mariadbBookRepo{Conn}
}

func (m *mariadbBookRepo) GetByID(ctx context.Context, id int64) (*models.Book, error) {
	return nil, nil
}
