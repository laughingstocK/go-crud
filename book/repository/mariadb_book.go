package repository

import (
	"database/sql"

	"github.com/laughingstocK/go-crud/book"
)

type mariadbBookRepo struct {
	Conn *sql.Conn
}

func NewMariadbBookRepo(Conn *sql.Conn) book.Repository {
	return &mariadbBookRepo{Conn}
}

// interface NewBookRepository()
