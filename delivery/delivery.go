package delivery

import (
	"gql/usecase"

	"github.com/graphql-go/graphql"
)

type resolver struct {
	productUsecase usecase.ProductUsecase
}

type Resolver interface {
	Search(params graphql.ResolveParams) (interface{}, error)
	AutoComplete(params graphql.ResolveParams) (interface{}, error)
}

func NewResolver(productUsecase usecase.ProductUsecase) Resolver {
	return &resolver{productUsecase}
}
