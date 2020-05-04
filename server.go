package main

import (
	"go-graphql/authentication"
	"go-graphql/database"
	"go-graphql/middelware"
	"go-graphql/route"
	"log"

	"github.com/labstack/echo"
)

func main() {
	dbconection, _ := database.ConnectDB()
	server := echo.New()
	r := server.Group("/graphql")
	//Database
	var data = database.NewDataBase(dbconection)
	//authentication
	var autheService = authentication.NewService(data)
	//----------------------------------------------//
	authentication.Route(r, autheService)

	middelware.ConfigMiddelware(r)

	r.GET("/", route.PlaygroundHandler)
	r.POST("/query", route.GraphQLHandler)
	log.Println(server.Start(":8080"))

}
