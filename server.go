package main

import (
	"gin-graphql-example/graph"
	"gin-graphql-example/storage"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Create the gqlgen server
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Storage: storage.NewLocalStorage(),
	}}))

	// Manually add the POST transport for JSON
	srv.AddTransport(transport.POST{})

	// If you want to support GET queries too
	srv.AddTransport(transport.GET{})

	// Define the GraphQL endpoint
	r.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	// Optional: Playaround/Sandbox UI
	r.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
	})

	r.Run()
}
