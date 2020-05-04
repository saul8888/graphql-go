package route

import (
	"go-graphql/graph"
	"go-graphql/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
)

func PlaygroundHandler(c echo.Context) error {
	playgroundHandler := playground.Handler("GraphQL playground", "/query")
	req := c.Request()
	res := c.Response()
	playgroundHandler.ServeHTTP(res, req)
	return nil
}

func GraphQLHandler(c echo.Context) error {
	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	req := c.Request()
	res := c.Response()
	graphqlHandler.ServeHTTP(res, req)
	return nil
}
