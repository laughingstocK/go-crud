package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/laughingstocK/go-crud/author"
	"github.com/laughingstocK/go-crud/models"
)

type mariadbAuthorRepo struct {
	Conn *sql.DB
}

func NewMariadbAuthorRepo(Conn *sql.DB) author.Repository {
	return &mariadbAuthorRepo{Conn}
}

func (m *mariadbAuthorRepo) GetByID(ctx context.Context, id int64) (*models.Author, error) {
	query := "SELECT id, name, createdAt, updatedAt FROM Author WHERE id = ?"
	row := m.Conn.QueryRowContext(ctx, query, id)

	var author models.Author
	err := row.Scan(&author.ID, &author.Name, &author.CreatedAt, &author.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("author not found")
		}
		log.Printf("Error retrieving author with ID %d: %v", id, err)
		return nil, fmt.Errorf("failed to retrieve author")
	}

	return &author, nil
}

func (m *mariadbAuthorRepo) Create(ctx context.Context, name string) (*models.Author, error) {
	return nil, nil
}
