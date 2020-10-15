package main

import (
	"log"
	"main/schemas"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	cepSchema := generateCEPSchema()

	gqlHandler := handler.New(&handler.Config{
		Schema:     &cepSchema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	router := mux.NewRouter().StrictSlash(false)

	router.Handle("/graphql", gqlHandler)

	log.Fatal(http.ListenAndServe(":3001", router))
}

func generateCEPSchema() graphql.Schema {
	schema, err := schemas.GenerateCEPSchema()

	if err != nil {
		panic(err)
	}

	return schema
}
