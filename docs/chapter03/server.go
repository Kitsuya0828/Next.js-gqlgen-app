package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Kitsuya0828/Next.js-gqlgen-app/docs/chapter03/graph"
	"github.com/Kitsuya0828/Next.js-gqlgen-app/docs/chapter03/graph/services"
	"github.com/Kitsuya0828/Next.js-gqlgen-app/docs/chapter03/internal"
)


const (
	defaultPort = "8080"
	dbFile      = "./mygraphql.db"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s?_foreign_keys=on&parseTime=true", dbFile))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service := services.New(db)

	srv := handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{Resolvers: &graph.Resolver{
		Srv:     service,
	}}))
	
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}