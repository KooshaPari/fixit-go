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
{
	Id: 4,
    Title:       "Finish report",
    Description: "Complete the financial report for the meeting",
    DueDate:     "2024-09-28",
    Done:        false,
},
{
	Id: 5,
    Title:       "Call mom",
    Description: "Catch up with mom on the phone",
    DueDate:     "2024-09-27",
    Done:        false,
},
{
	Id: 6,
    Title:       "Grocery shopping",
    Description: "Buy vegetables and fruits for the week",
    DueDate:     "2024-09-29",
    Done:        false,
},
{
	Id: 7,
    Title:       "Car service",
    Description: "Take the car for its scheduled maintenance",
    DueDate:     "2024-10-01",
    Done:        false,
},
{
	Id: 8,
    Title:       "Attend meeting",
    Description: "Join the team meeting at 3 PM",
    DueDate:     "2024-09-30",
    Done:        true,
},
{
	Id: 9,
    Title:       "Water plants",
    Description: "Water all indoor and outdoor plants",
    DueDate:     "2024-09-26",
    Done:        false,
},
{
	Id: 10,
    Title:       "Clean kitchen",
    Description: "Wipe counters, sweep floors, and clean the sink",
    DueDate:     "2024-09-27",
    Done:        false,
},
{
	Id: 11,
    Title:       "Workout",
    Description: "Complete a 45-minute cardio workout",
    DueDate:     "2024-09-26",
    Done:        false,
},
{
	Id: 12,
    Title:       "Read book",
    Description: "Finish reading 50 pages of the current novel",
    DueDate:     "2024-09-28",
    Done:        false,
},
{
	Id: 13,
    Title:       "Prepare presentation",
    Description: "Draft the slides for Monday's presentation",
    DueDate:     "2024-09-30",
    Done:        false,
},
{
	Id: 14,
    Title:       "Pay electricity bill",
    Description: "Pay the bill before the due date",
    DueDate:     "2024-09-28",
    Done:        true,
},
{
	Id: 15,
    Title:       "Bake a cake",
    Description: "Prepare a cake for the weekend get-together",
    DueDate:     "2024-09-29",
    Done:        false,
},
{
	Id: 16,
    Title:       "Plan vacation",
    Description: "Research hotels and activities for the trip",
    DueDate:     "2024-10-05",
    Done:        false,
},
{
	Id: 17,
    Title:       "Visit the doctor",
    Description: "Annual check-up appointment at 10 AM",
    DueDate:     "2024-09-30",
    Done:        false,
},
{
	Id: 18,
    Title:       "Take out trash",
    Description: "Take the trash bins to the curb",
    DueDate:     "2024-09-24",
    Done:        true,
},
{
	Id: 19,
    Title:       "Organize closet",
    Description: "Rearrange and declutter the closet",
    DueDate:     "2024-09-28",
    Done:        false,
},
{
	Id: 20,
    Title:       "Prepare dinner",
    Description: "Cook a healthy meal for family dinner",
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
