package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Kitsuya0828/Next.js-gqlgen-app/backend/graph"
	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const defaultPort = "8080"

func main() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(aws.connect.psdb.cloud)/%s?tls=true",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DATABASE_NAME"),
	)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if db == nil {
		log.Fatal(err)
	}
	defer func() {
		if db != nil {
			if sqlDB, err := db.DB(); err != nil {
				log.Fatal(err)
			} else {
				if err := sqlDB.Close(); err != nil {
					log.Fatal(err)
				}
			}
		}
	}()
	db.Logger = db.Logger.LogMode(logger.Info)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: *db}}))
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
