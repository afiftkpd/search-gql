package delivery

import (
	"context"
	"fmt"

	"github.com/graphql-go/graphql"
)

func (r *resolver) Search(params graphql.ResolveParams) (interface{}, error) {
	keyword := ""
	if keywordFromClients, ok := params.Args["keyword"].(string); ok {
		keyword = keywordFromClients
	}

	page := 0
	if pageFromClients, ok := params.Args["page"].(int); ok {
		page = pageFromClients
	}

	limit := 0
	if pageFromClients, ok := params.Args["limit"].(int); ok {
		limit = pageFromClients
	}

	fmt.Println("heree")
	fmt.Println(params.Args)

	results, err := r.productUsecase.Search(context.Background(), keyword, page, limit)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *resolver) AutoComplete(params graphql.ResolveParams) (interface{}, error) {
	keyword := ""
	if keywordFromClients, ok := params.Args["keyword"].(string); ok {
		keyword = keywordFromClients
	}

	results, err := r.productUsecase.AutoComplete(context.Background(), keyword)
	if err != nil {
		return nil, err
	}

	return results, nil
}
