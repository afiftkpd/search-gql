package delivery

import (
	"github.com/graphql-go/graphql"
)

// NewSchema initializes Schema struct which takes resolver as the argument.
func NewSchema(productResolver Resolver) Schema {
	return Schema{
		productResolver: productResolver,
	}
}

// Schema is struct which has method for Query and Mutation. Please init this struct using constructor function.
type Schema struct {
	productResolver Resolver
}

var SearchResultGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SearchResult",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Int,
			},
			"rating": &graphql.Field{
				Type: graphql.Int,
			},
			"image_url": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"stock": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var AutoCompleteResultGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AutoCompleteResult",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// Query initializes config schema query for graphql server.
func (s Schema) Query() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"Search": &graphql.Field{
				Type:        graphql.NewList(SearchResultGraphQL),
				Description: "Search Product",
				Args: graphql.FieldConfigArgument{
					"keyword": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"page": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: s.productResolver.Search,
			},
			"AutoComplete": &graphql.Field{
				Type:        graphql.NewList(AutoCompleteResultGraphQL),
				Description: "Autocomplete",
				Args: graphql.FieldConfigArgument{
					"keyword": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: s.productResolver.AutoComplete,
			},
		},
	}

	return graphql.NewObject(objectConfig)
}
