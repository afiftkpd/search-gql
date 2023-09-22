package main

import (
	"log"
	"net/http"

	"gql/delivery"
	repoHttp "gql/repository/http"
	"gql/usecase"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {

	// Initialize Repo, Usecase, & Schema Layer
	repo := repoHttp.NewProductRepository()
	usecase := usecase.NewProductUsecase(repo)
	schema := delivery.NewSchema(delivery.NewResolver(usecase))

	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: schema.Query(),
		//Mutation: schema.Mutation(),
	})

	if err != nil {
		log.Fatal(err)
	}

	gqlHandler := handler.New(&handler.Config{
		Schema:     &graphqlSchema,
		GraphiQL:   false,
		Pretty:     true,
		Playground: true,
	})

	http.Handle("/graphql", CorsMiddleware(gqlHandler))
	log.Println("Starting Service-GQL at port 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

// CORS Handler
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
