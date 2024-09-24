package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Title  string `json:"title"`
	description  string `json:"description"`
	dueDate string `json:"dueDate"`
	Done bool `json:"done"`
}

var todos = []Todo{
  {
    Title:       "Buy milk",
    description: "Buy 2 gallons of milk from the store",
    dueDate:     "2024-09-25",
    Done:        false,
  },
  {
    Title:       "Walk the dog",
    description: "Take the dog for a 30-minute walk",
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

	jsonBytes, err := json.Marshal((todolist))
	if err != nil {
    c.JSON(500, gin.H{"error": err.Error()})
    return
}
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