package template

import (
	"bytes"
	"fmt"
	"text/template"
)

func GetMainTemplate(moduleName string) (string, error) {
	mainConfig := MainConfig{
		ModuleName: moduleName,
	}

	mainTemplate, err := template.New("mainTemplate").Parse(mainTemplate)
	if err != nil {
		return "", fmt.Errorf("err parse template mainTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	if err = mainTemplate.Execute(&templateBuf, mainConfig); err != nil {
		return "", fmt.Errorf("err create template: %v", err)
	}
	return templateBuf.String(), nil
}

type MainConfig struct {
	ModuleName string
}

var mainTemplate = `package main

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
