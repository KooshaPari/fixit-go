package main

import (
	"github.com/gin-gonic/gin"
)
type Todo struct{
	Title string `json:"title"`;
	Description string
}
func main() {
	r := gin.Default()
	r.GET("/todos", getTodos)
	r.POST("/todos", createTodo)
	r.GET("/todos/:id", getTodo)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)
	r.Run(":8080")
}

func getTodos(c *gin.Context) {
	// TO DO: implement getting all todos
}

func createTodo(c *gin.Context) {
	// TO DO: implement creating a new todo
}

func getTodo(c *gin.Context) {
	// TO DO: implement getting a single todo
}

func updateTodo(c *gin.Context) {
	// TO DO: implement updating a todo
}

func deleteTodo(c *gin.Context) {
	// TO DO: implement deleting a todo
}