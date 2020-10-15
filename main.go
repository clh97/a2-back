package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/friendsofgo/graphiql"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home! :)")
}

func main() {
	graphiqlHandler := initializeGraphiQL()

	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/", homeLink)

	router.Handle("/graphiql", graphiqlHandler)

	log.Fatal(http.ListenAndServe(":3001", router))
}

func initializeGraphiQL() *graphiql.Handler {
	handler, err := graphiql.NewGraphiqlHandler("/graphql")

	if err != nil {
		panic(err)
	}

	return handler
}
