package author

import (
	"context"

	"github.com/laughingstocK/go-crud/models"
)

type Repository interface {
	GetByID(ctx context.Context, id int64) (*models.Author, error)
	Create(ctx context.Context, name string) (*models.Author, error)
}
