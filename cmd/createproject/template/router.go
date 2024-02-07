package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GetRouterTemplate(domain string) (string, error) {
	routerConfig := NewRouterConfig(domain)

	routerTemplate, err := template.New("routerTemplate").Parse(routerTemplate)
	if err != nil {
		return "", fmt.Errorf("err parse template routerTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	if err = routerTemplate.Execute(&templateBuf, routerConfig); err != nil {
		return "", fmt.Errorf("err create template: %v", err)
	}

	return templateBuf.String(), nil
}

type RouterConfig struct {
	DomainPascalCase string
	DomainCamelCase  string
}

func NewRouterConfig(domain string) RouterConfig {
	return RouterConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
	}
}

var routerTemplate = `package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(ginEngine *gin.Engine, db *gorm.DB) {
	{{.DomainCamelCase}}Controller := get{{.DomainPascalCase}}Controller(db)

	ginEngine.GET("", {{.DomainCamelCase}}Controller.Qux)
}
`
