package templ

var CMDMain = `package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"{{.ModuleName}}/database"
	"{{.ModuleName}}/internal/router"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	db, err := database.InitializeDB()
	if err != nil {
		log.Fatal("err initialize db")
	}

	ginEngine := gin.Default()

	router.Register(ginEngine, db)

	if err := ginEngine.Run(); err != nil {
		log.Fatal(err)
	}
}
`
