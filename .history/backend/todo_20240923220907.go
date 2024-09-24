package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Title  string `json:"title"`
	Description  string `json:"description"`
	DueDate string `json:"dueDate"`
	Done bool `json:"done"`
}

var todos = []Todo{
  {
    Title:       "Buy milk",
    Description: "Buy 2 gallons of milk from the store",
    DueDate:     "2024-09-25",
    Done:        false,
  },
  {
    Title:       "Walk the dog",
    Description: "Take the dog for a 30-minute walk",
    DueDate:     "2024-09-24",
    Done:        true,
  },
  {
    Title:       "Do laundry",
    Description: "Wash, dry, and fold the laundry",
    DueDate:     "2024-09-26",
    Done:        false,
  },
}

func GetTodos(c *gin.Context){
	todolist := todos
	
	c.JSON(200,jsonBytes)
}
func createTodo(/*TODO*/) {
	//TODO
}

func getTodo(/*TODO*/) {
	//TODO
}

func updateTodo(/*TODO*/) {
	/*TODO*/
}
func deleteTodo(/*TODO*/) {
	/*TODO*/
}