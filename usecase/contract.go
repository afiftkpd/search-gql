package usecase

import (
	"context"
	"gql/models"
)

type ProductUsecase interface {
	Search(ctx context.Context, keyword string, page int, limit int) ([]models.Product, error)
	AutoComplete(ctx context.Context, keyword string) ([]models.AutoComplete, error)
}
