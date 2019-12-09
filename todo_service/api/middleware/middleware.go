package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"test/gin-gonic/todo_service/constant"
	"test/gin-gonic/todo_service/dal"
)

func CockroachDbMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		dbConn, err := dal.Connect()
		if err != nil {
			fmt.Println("database connection failed:", err)
			log.Fatal(err)
		}

		c.Set(constant.DBConnCtxKey, dbConn)
		c.Next()
	}
}
