package main

import (
	"fmt"
	"log"
	"main/schemas"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home! :)")
}

func main() {
	cepSchema := generateCEPSchema()

	gqlHandler := handler.New(&handler.Config{
		Schema:     &cepSchema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/", homeLink)
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
