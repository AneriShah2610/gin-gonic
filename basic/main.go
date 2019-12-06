package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", GetHomePage)
	router.POST("/", PostHomePage)
	router.GET("/query", QueryString) // query?name=abc&age=01
	router.GET("/path/:name/:age", PathParameters) // path/abc/24
	router.POST("/body", PostHomePageWithBody) // path/abc/24
	router.Run(":8080")
}

func GetHomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world!!",
	})
}

func PostHomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Post home page",
	})
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}


func PathParameters(c *gin.Context){
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func PostHomePageWithBody(c *gin.Context)  {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil{
		fmt.Println("Error", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": string(value),
	})
}