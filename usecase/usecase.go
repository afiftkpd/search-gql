package usecase

import (
	"context"
	"gql/models"
	"gql/repository/http"
)

type productUsecase struct {
	Repo http.ProductRepository
}

func NewProductUsecase(repo http.ProductRepository) ProductUsecase {
	return &productUsecase{repo}
}

func (p *productUsecase) Search(ctx context.Context, keyword string, page int, limit int) ([]models.Product, error) {
	return p.Repo.Search(ctx, keyword, page, limit)
}

func (p *productUsecase) AutoComplete(ctx context.Context, keyword string) ([]models.AutoComplete, error) {
	return p.Repo.AutoComplete(ctx, keyword)
}
