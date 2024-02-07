package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GetInjectionTemplate(domain string, moduleName string) (string, error) {
	injectionConfig := NewInjectionConfig(domain)
	injectionConfig.ModuleName = moduleName

	injectionTemplate, err := template.New("injectionTemplate").Parse(injectionTemplate)
	if err != nil {
		return "", fmt.Errorf("err parse template injectionTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	if err = injectionTemplate.Execute(&templateBuf, injectionConfig); err != nil {
		return "", fmt.Errorf("err create template: %v", err)
	}

	return templateBuf.String(), nil
}

type InjectionConfig struct {
	DomainPascalCase string
	DomainCamelCase  string
	ModuleName       string
}

func NewInjectionConfig(domain string) InjectionConfig {
	return InjectionConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
	}
}

var injectionTemplate = `package router

import (
	"gorm.io/gorm"
	"{{.ModuleName}}/internal/controller/http"
	"{{.ModuleName}}/internal/repository"
	"{{.ModuleName}}/internal/service"
)

func get{{.DomainPascalCase}}Controller(db *gorm.DB) *http.{{.DomainPascalCase}}Controller {
	{{.DomainCamelCase}}Repository := repository.New{{.DomainPascalCase}}Repository(db)
	{{.DomainCamelCase}}Service := service.New{{.DomainPascalCase}}Service({{.DomainCamelCase}}Repository)
	{{.DomainCamelCase}}Controller := http.New{{.DomainPascalCase}}Controller({{.DomainCamelCase}}Service)
	return {{.DomainCamelCase}}Controller
}
`
