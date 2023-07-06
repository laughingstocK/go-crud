package book

import (
	"context"

	"github.com/laughingstocK/go-crud/models"
)

type Repository interface {
	GetByID(ctx context.Context, id int64) (*models.Book, error)
}
