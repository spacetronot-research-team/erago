package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GetControllerTemplate(domain string, moduleName string, varErr1 string) (string, error) {
	controllerConfig := NewControllerConfig(domain, varErr1)
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
	VarErr1          string
}

func NewControllerConfig(domain string, varErr1 string) ControllerConfig {
	return ControllerConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		DomainShort:      getDomainShort(domain),
		VarErr1:          varErr1,
	}
}

var controllerTemplate = `package http

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"{{.ModuleName}}/internal/service"
)

var (
	Err{{.VarErr1}} = errors.New("err jklasjd")
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
		err = errors.Join(err, Err{{.VarErr1}})
		log.Println(err)
		return
	}
	log.Println("^.^")
}
`
