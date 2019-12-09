package dal

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"test/gin-gonic/todo_service/config"
)

var instance *sql.DB
var once sync.Once

func Connect() (*sql.DB, error) {
	once.Do(func() {
		connectionString := fmt.Sprintf("postgresql://%s@%s:%s/%s?sslmode=disable", config.Conf.Cockroach.User, config.Conf.Cockroach.Host, config.Conf.Cockroach.Port, config.Conf.Cockroach.DbName)

		db, err := sql.Open("postgres", connectionString)
		if err != nil {
			log.Println("error while initializing database", err)
			return
		}
		fmt.Println("Database successfully initialized")
		instance = db
	})
	return instance, nil
}
