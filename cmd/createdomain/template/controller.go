package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GetControllerTemplate(domain string, moduleName string) (string, error) {
	controllerConfig := NewControllerConfig(domain)
	controllerConfig.ModuleName = moduleName

	controllerTemplate, err := template.New("controllerTemplate").Parse(controllerTemplate)
	if err != nil {
		return "", fmt.Errorf("err parse template controllerTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	if err = controllerTemplate.Execute(&templateBuf, controllerConfig); err != nil {
		return "", fmt.Errorf("err create template: %v", err)
	}

	return templateBuf.String(), nil
}

type ControllerConfig struct {
	DomainPascalCase string
	DomainCamelCase  string
	DomainShort      string
	ModuleName       string
}

func NewControllerConfig(domain string) ControllerConfig {
	return ControllerConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		DomainShort:      getDomainShort(domain),
	}
}

var controllerTemplate = `package http

import (
	"log"

	"github.com/gin-gonic/gin"
	"{{.ModuleName}}/internal/service"
)

type {{.DomainPascalCase}}Controller struct {
	{{.DomainCamelCase}}Service service.{{.DomainPascalCase}}
}

func New{{.DomainPascalCase}}Controller({{.DomainCamelCase}}Service service.{{.DomainPascalCase}}) *{{.DomainPascalCase}}Controller {
	return &{{.DomainPascalCase}}Controller{
		{{.DomainCamelCase}}Service: {{.DomainCamelCase}}Service,
	}
}

// Qux babibu.
func ({{.DomainShort}}c *{{.DomainPascalCase}}Controller) Qux(ctx *gin.Context) {
	if err := {{.DomainShort}}c.{{.DomainCamelCase}}Service.Bar(ctx); err != nil {
		log.Println(err)
		return
	}
	log.Println("^.^")
}
`
