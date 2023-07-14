package author

import (
	"context"

	"github.com/laughingstocK/go-crud/models"
)

type Usecase interface {
	GetByID(ctx context.Context, id int64) (*models.Author, error)
	Create(context.Context, *models.Author) (*models.Author, error)
}
