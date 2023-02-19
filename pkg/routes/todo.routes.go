package routes

import (
	"github.com/GaurKS/todo-app/pkg/services"
	"github.com/gin-gonic/gin"
)

func TodoRouter(r *gin.RouterGroup) {
	r.POST("/create", services.CreateTodo)
	r.GET("/read/all", services.ReadAllTodos)
	r.GET("/read/:id", services.ReadTodoById)
	r.PATCH("/update/:id", services.UpdateTodoById)
	r.DELETE("/delete/:id", services.DeleteTodoById)
	r.POST("/parse/csv", services.ParseCsv)
	r.GET("/health", services.HealthCheck)
}