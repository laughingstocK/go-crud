package usecase

import (
	"context"

	"github.com/laughingstocK/go-crud/author"
	"github.com/laughingstocK/go-crud/models"
)

type authorUsecase struct {
	mariadbAuthorRepo author.Repository
	grpcAuthorRepo    author.Repository
}

func NewAuthorUsecase(mariadbAuthorRepo, grpcAuthorRepo author.Repository) author.Usecase {
	return &authorUsecase{mariadbAuthorRepo, grpcAuthorRepo}
}

func (a *authorUsecase) GetByID(ctx context.Context, id int64) (*models.Author, error) {
	resAuthor, err := a.mariadbAuthorRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return resAuthor, nil
}

func (a *authorUsecase) Create(ctx context.Context, author *models.Author) (*models.Author, error) {
	resAuthor, err := a.grpcAuthorRepo.Create(ctx, author.Name)
	if err != nil {
		return nil, err
	}
	return resAuthor, nil
}
