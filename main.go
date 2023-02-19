package main

import (
	"log"
	"os"

	"github.com/GaurKS/todo-app/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Setting variables for app env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	router.Use(gin.Logger())
	r := router.Group("/api")
	routes.TodoRouter(r.Group("/todo"))

	router.Run(":" + os.Getenv("PORT"))
}