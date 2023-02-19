package main

import (
	"log"
	"os"

	"github.com/GaurKS/todo-app/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// Setting variables for app env
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	router := gin.Default()
	router.Use(gin.Logger())
	r := router.Group("/api")
	routes.TodoRouter(r.Group("/todo"))

	port := os.Getenv("PORT")
	if port == "" {
	 	port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
	 	log.Panicf("error: %s", err)
	}
}