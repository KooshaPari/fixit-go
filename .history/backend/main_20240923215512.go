package main

import (
	"github.com/gin-gonic/gin"
	 "github.com/gin-contrib/cors"
router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8080"}, // Frontend origin
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

)

func main() {
	r := gin.Default()
	r.GET("/api/todos", GetTodos)
	/*r.POST("/todos", createTodo)
	r.GET("/todos/:id", getTodo)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)*/
	r.Run(":8080")
}
