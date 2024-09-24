package main

import (
	"net/http"

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
  {
	Id: 1,
    Title:       "Buy milk",
    Description: "Buy 2 gallons of milk from the store",
    DueDate:     "2024-09-25",
    Done:        false,
  },
  {
	Id: 2,
    Title:       "Walk the dog",
    Description: "Take the dog for a 30-minute walk",
    DueDate:     "2024-09-24",
    Done:        true,
  },
  {
	Id: 3,
    Title:       "Do laundry",
    Description: "Wash, dry, and fold the laundry",
    DueDate:     "2024-09-26",
    Done:        false,
  },
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

func getTodo(/*TODO*/) {
	//TODO
}

func updateTodo(/*TODO*/) {
	/*TODO*/
}
func deleteTodo(c *gin.Context) {
	for i, todo:= range todos{}
}