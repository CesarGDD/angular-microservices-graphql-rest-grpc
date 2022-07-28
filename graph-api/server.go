package main

import (
	"cesargdd/graph-api/directives"
	"cesargdd/graph-api/graph"
	"cesargdd/graph-api/graph/generated"
	"cesargdd/graph-api/middlewares"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
		AllowedHeaders:   []string{"*"},
	}).Handler)
	router.Use(middlewares.AuthMiddleware)

	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Auth = directives.Auth
	srv := handler.New(generated.NewExecutableSchema(c))
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", (srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
