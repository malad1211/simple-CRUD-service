package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"inspiredlab/handler"
	"inspiredlab/repository"
	"inspiredlab/service"
	"log"
	"os"
)

var appEnv = "DEV"

func init() {
	env := ".env"
	profile := "default"
	if os.Getenv("profile") != "" {
		profile = os.Getenv("profile")
	}
	fmt.Println("Run with profile: ", profile)
	if profile != "default" {
		env = fmt.Sprintf("%v.%v", env, profile)
	}
	if err := godotenv.Load(env); err != nil {
		log.Println("no env option")
	}
}

func main() {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_CONN")))
	if err != nil {
		log.Panic(err)
	}

	if err != nil {
		log.Panic(err)
	}

	host := ""
	appPort := os.Getenv("APP_PORT")

	r := gin.Default()

	//repo
	newsRepo := repository.NewNewsRepository(db)

	//service
	newsService := service.NewNewsService(newsRepo)

	//handler
	newsGroup := r.Group("news")
	handler.NewNewsHandler(newsGroup, newsService)

	err = r.Run(host + ":" + appPort)
	if err != nil {
		log.Panic(err)
	}
}
