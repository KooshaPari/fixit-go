package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8080"}, // Frontend origin
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))
	r.GET("/api/todos", GetTodos)
	r.POST("/api/todos", createTodo)
	r.DELETE("/api/todos/:id", deleteTodo)
	r.GET("/api/todos/:id", getTodo)
	r.PUT("/api/todos/:id", updateTodo)
	
	r.Run(":8081")
}
