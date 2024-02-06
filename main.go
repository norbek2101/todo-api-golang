package main

import (
	"todo-api-golang/db"
	"todo-api-golang/docs"
	"todo-api-golang/routes"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	
)

// @title Todo API
// @version 1.0
// @description This is a simple To-Do API using Golang and PostgreSQL
// @host localhost:8080
// @BasePath /
func main() {
	database.Connect()

	router := routes.SetupRouter()

	// // programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Swagger documentation
	// url := ginSwagger.URL("http://localhost:8080/docs/swagger.json") // Point to your generated doc.json
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
