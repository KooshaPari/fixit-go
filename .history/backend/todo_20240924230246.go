package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id int `json:"id"`
	Title  string `json:"title"`
	Description  string `json:"description"`
	DueDate string `json:"dueDate"`
	Done bool `json:"done"`
}

var todos = []Todo{
    {"id": 1, "title": "Test Todo 1", "completed": false},
    {"id": 2, "title": "Test Todo 2", "completed": false},
    {"id": 3, "title": "Test Todo 3", "completed": false},
    {"id": 4, "title": "Test Todo 4", "completed": false},
    {"id": 5, "title": "Test Todo 5", "completed": false},
    {"id": 6, "title": "Test Todo 6", "completed": false},
    {"id": 7, "title": "Test Todo 7", "completed": false},
    {"id": 8, "title": "Test Todo 8", "completed": false},
    {"id": 9, "title": "Test Todo 9", "completed": false},
    {"id": 10, "title": "Test Todo 10", "completed": false},
}

func GetTodos(c *gin.Context){
	todolist := todos
	
	c.JSON(200,todolist)
}
func createTodo(c *gin.Context) {
	var newTodo Todo
	err := c.BindJSON(&newTodo)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
	newTodo.Id = len(todos)+1;
	
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)

}

func getTodo(c *gin.Context) {
	id := c.Param("id");
	for i, todo := range todos {
		if strconv.Itoa(todo.Id) == id {
			curTodo := todos[i]
			c.JSON(http.StatusOK, curTodo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo item not found"})
}

func editTodo(c *gin.Context) {
	fmt.Println("Request Received For Edit")
	id := c.Param("id");
	var updatedTodo Todo
	c.BindJSON(&updatedTodo)
	for i, todo := range todos {
		if strconv.Itoa(todo.Id) == id{
			todos[i] = updatedTodo;
			c.JSON(http.StatusOK, todo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo item not found"})
}
func deleteTodo(c *gin.Context) {
	id := c.Param("id");
	for i, todo := range todos {
		if strconv.Itoa(todo.Id) == id {
			deletedTodo := todos[i]
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, deletedTodo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo item not found"})
}