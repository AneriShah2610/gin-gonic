package handler

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/gin-gonic/todo_service/constant"
	"test/gin-gonic/todo_service/models"
	"time"
)

func CreateTodo(c *gin.Context) {
	db := c.Keys[constant.DBConnCtxKey].(*sql.DB)
	var todo models.RequestTodoModel
	var err error
	err = c.BindJSON(&todo)
	if err != nil {
		fmt.Println("Error at binding data:", err)
		return
	}
	_, err = db.Exec("INSERT INTO todo.todo (title, created_by, created_at) values ($1, $2, $3)", todo.Title, todo.CreatedBy, time.Now().UTC())
	if err != nil {
		fmt.Println("Error while inserting:", err)
		return
	}
	c.JSON(http.StatusOK, "Todo successfully inserted")
}

func FetchAllTodos(c *gin.Context) {
	db := c.Keys[constant.DBConnCtxKey].(*sql.DB)
	var todos []models.ResponseTodoModel

	rows, err := db.Query("SELECT id, title, todo_status, created_by, created_at FROM todo")
	if err != nil {
		fmt.Println("Error while fetching:", err)
		return
	}

	for rows.Next() {
		var todo models.ResponseTodoModel
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Status, &todo.CreatedBy, &todo.CreatedAt)
		if err != nil {
			fmt.Println("Error while scanning:", err)
			return
		}
		todos = append(todos, todo)
	}
	c.JSON(http.StatusOK, todos)
}

func FetchParticularUsersTodos(c *gin.Context) {
	db := c.Keys[constant.DBConnCtxKey].(*sql.DB)
	createdBy := c.Param("name")
	todos := []models.ResponseTodoModel{}
	rows, err := db.Query("SELECT id, title, todo_status, created_at FROM todo WHERE created_by = $1", createdBy)
	if err != nil {
		fmt.Println("Error while fetching:", err)
		return
	}

	for rows.Next() {
		var todo models.ResponseTodoModel
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Status, &todo.CreatedAt)
		if err != nil {
			fmt.Println("Error while scanning:", err)
			return
		}
		todos = append(todos, todo)
	}
	c.JSON(http.StatusOK, todos)
}

func DeleteTodo(c *gin.Context) {
	db := c.Keys[constant.DBConnCtxKey].(*sql.DB)
	title := c.Param("todo_name")
	userName := c.Param("name")

	_, err := db.Exec("DELETE FROM todo WHERE title = $1 AND created_by = $2", title, userName)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	c.JSON(http.StatusOK, "Todo deleted successfully")
}

func UpdateTodoData(c *gin.Context) {
	db := c.Keys[constant.DBConnCtxKey].(*sql.DB)
	oldTitle := c.Param("todo_name")
	userName := c.Param("name")
	var todo models.ResponseTodoModel
	row := db.QueryRow("SELECT title, todo_status FROM todo WHERE title = $1 AND created_by = $2", oldTitle, userName)
	err := row.Scan(&todo.Title, &todo.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, "Yo have no such todo")
			return
		}
		c.JSON(http.StatusOK, err)
		fmt.Println("Error while fetching data..", err)
		return
	}

	var updatedTodoData models.UpdateTodoModel
	err = c.BindJSON(&updatedTodoData)
	if err != nil {
		c.JSON(http.StatusOK, err)
		fmt.Println("Error at binding data:", err)
		return
	}

	if updatedTodoData.NewTitle == nil {
		updatedTodoData.NewTitle = &todo.Title
	}
	if updatedTodoData.Status == nil {
		updatedTodoData.Status = &todo.Status
	}

	_, err = db.Exec("UPDATE todo SET (title, todo_status) = ($1, $2)  WHERE title = $3 AND created_by = $4", updatedTodoData.NewTitle, updatedTodoData.Status, oldTitle, userName)
	if err != nil {
		c.JSON(http.StatusOK, err)
		fmt.Println("Error while updating todo data:", err)
		return
	}
	c.JSON(http.StatusOK, "Todo data updated successfully!!")
}
