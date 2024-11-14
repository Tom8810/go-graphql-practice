package main

import (
	"fmt"
	"gorm_practice1/graph"
	database "gorm_practice1/internal/db"
	getenv "gorm_practice1/internal/env"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	if err := getenv.Get("./internal/env/.env"); err != nil {
		fmt.Printf("cannnot read environmental variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
