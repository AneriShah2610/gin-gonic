package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"test/gin-gonic/todo_service/api/handler"
	"test/gin-gonic/todo_service/api/middleware"
	"test/gin-gonic/todo_service/config"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		fmt.Println("conf loading failed:", err)
		log.Fatal(err)
	}

	router := gin.Default()

	router.Use(middleware.CockroachDbMiddleware())
	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", handler.CreateTodo)
		v1.GET("/", handler.FetchAllTodos)
		v1.GET("/user/:name", handler.FetchParticularUsersTodos)
		v1.DELETE("/:todo_name/user/:name", handler.DeleteTodo)
		v1.PUT("/:todo_name/user/:name", handler.UpdateTodoData)
	}
	log.Printf("Connect to http://%s:%s", conf.Server.Host, conf.Server.Port)
	router.Run(conf.Server.Port)
}
