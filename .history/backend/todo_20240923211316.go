package main
import ( "encoding/json"
"fmt")

type Todo struct {
	Title  string `json:"title"`
	Description  string `json:"Description"`
	dueDate string `json:"dueDate"`
	Done bool `json:"done"`
}

var todos []Todo

func getTodos(/**/){
	//TODO
}
func createTodo(/*TODO*/) {
	//TODO
}

func getTodo(/*TODO*/) {
	//TODO
}

func updateTodo(/*TODO*/) {
	/*TODO*/

func deleteTodo(/*TODO*/) {
	/*TODO*/
}