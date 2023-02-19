package services

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/GaurKS/todo-app/pkg/dtos"
	"github.com/gin-gonic/gin"
)

var todos = []dtos.Todo{
	{ID: "1", Title: "Todo 1", TodoStatus: "In Progress", Description: "Todo 1 Description", CreatedBy: "Gaurav"},
	{ID: "2", Title: "Todo 2", TodoStatus: "Completed", Description: "Todo 2 Description", CreatedBy: "Alex"},
	{ID: "3", Title: "Todo 3", TodoStatus: "In Progress", Description: "Todo 3 Description", CreatedBy: "David"},
}

func GenerateId(strlen int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABC0123456789"
	result := make([]byte, strlen)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

func CreateTodo(c *gin.Context) {
	todo := dtos.Todo{}

	if err := c.BindJSON(&todo); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		fmt.Println(err.Error())
		c.Abort()
		return
	}

	todo.ID = GenerateId(5)
	todos = append(todos, todo)

	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"data": todo,
		},
	)
}

func ReadAllTodos(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"data": todos,
		},
	)
}

func ReadTodoById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid ID",
			},
		)
		c.Abort()
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			c.IndentedJSON(
				http.StatusOK,
				gin.H{
					"data": todo,
				},
			)
			return
		}
	}

	c.IndentedJSON(
		http.StatusNotFound,
		gin.H{
			"error": "Todo not found",
		},
	)
}

func UpdateTodoById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid ID",
			},
		)
		c.Abort()
		return
	}

	for index, todo := range todos {
		if todo.ID == id {
			updatedTodo := dtos.Todo{}

			if err := c.BindJSON(&updatedTodo); err != nil {
				c.IndentedJSON(
					http.StatusBadRequest,
					gin.H{
						"error": err.Error(),
					},
				)
				fmt.Println(err.Error())
				c.Abort()
				return
			}

			todo.Title = updatedTodo.Title
			todo.TodoStatus = updatedTodo.TodoStatus
			todo.Description = updatedTodo.Description
			todo.CreatedBy = updatedTodo.CreatedBy

			todos[index] = todo

			c.IndentedJSON(
				http.StatusOK,
				gin.H{
					"data": todo,
				},
			)
			return
		}
	}

	c.IndentedJSON(
		http.StatusNotFound,
		gin.H{
			"error": "Todo not found",
		},
	)
}

func DeleteTodoById(c *gin.Context) { 
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid ID",
			},
		)
		c.Abort()
		return
	}

	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			c.IndentedJSON(
				http.StatusOK,
				gin.H{
					"message": "Todo deleted successfully",
				},
			)
			return
		}
	}

	c.IndentedJSON(
		http.StatusNotFound,
		gin.H{
			"error": "Todo not found",
		},
	)
}

func ParseCsv(c *gin.Context) {
	file, err := c.FormFile("csv")
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest, 
			gin.H {
				"error": "File not found",
			},
		)
		c.Abort()
		return
	}

	f, err := file.Open()
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest, 
			gin.H {
				"error": "Failed to open file",
			},
		)
		c.Abort()
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	var todos []dtos.Todo

	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			c.IndentedJSON(
				http.StatusBadRequest, 
				gin.H {
					"error": "Failed to parse CSV file",
				},
			)
			c.Abort()
			return
		}

		todo := dtos.Todo{
			Title: row[0],
			TodoStatus: row[1],
			Description: row[2],
			CreatedBy: row[3],
		}

		todos = append(todos, todo)
	}

	var builder strings.Builder
	builder.WriteString("+--------------------------+------------------+------------------------------------+--------------------+\n")
	for _, record := range todos {
			builder.WriteString(fmt.Sprintf("| %-24s | %-16s | %-34s | %-18s |\n", 
				record.Title, 
				record.TodoStatus, 
				record.Description, 
				record.CreatedBy,
			))
	}
	builder.WriteString("+--------------------------+------------------+------------------------------------+--------------------+\n")

	// Return the formatted table
	c.String(http.StatusOK, builder.String())
}

func HealthCheck(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{
			"message": "Server is up and running...🚀",
		},
	)	
}