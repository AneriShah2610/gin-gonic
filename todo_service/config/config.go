package config

import (
	"fmt"
	environment "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path"
	er "test/gin-gonic/todo_service/error"
)

type Service struct {
	Server    Server
	Cockroach Cockroach
}

type Server struct {
	Host string `env:"HOST"`
	Port string `env:"PORT"`
}

type Cockroach struct {
	Host     string `env:"DBHOST"`
	Port     string `env:"DBPORT"`
	User     string `env:"DBUSER"`
	DbName   string `env:"DBNAME"`
	Password string `env:"DBPASSWORD"`
}

var Conf Service
var isLoaded = false
var Env string

func Load() (Service, error) {
	if isLoaded == true {
		return Conf, nil
	}

	gopath := os.Getenv("GOPATH")
	applicationPath := path.Join(gopath, "/src/test/gin-gonic/todo_service/.config")
	dir := fmt.Sprintf("%s/.env", applicationPath)
	err := godotenv.Load(dir)
	if err != nil {
		log.Fatal(err)
	}

	Env = getEnv("ENV", "")
	if Env != "dev" && Env != "prod" && Env != "stage" && Env != "test" {
		log.Fatal(er.InvalidEnvironment)
	}

	valueStr := getEnv("DBPORT", "")
	Conf.Cockroach.Port = valueStr

	_, err = environment.UnmarshalFromEnviron(&Conf)
	if err != nil {
		return Service{}, err
	}
	isLoaded = true
	return Conf, err
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
