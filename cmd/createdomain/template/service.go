package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GetServiceTemplate(domain string, moduleName string) (string, error) {
	serviceConfig := NewServiceConfig(domain)
	serviceConfig.ModuleName = moduleName

	serviceTemplate, err := template.New("serviceTemplate").Parse(serviceTemplate)
	if err != nil {
		return "", fmt.Errorf("err parse template serviceTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	if err = serviceTemplate.Execute(&templateBuf, serviceConfig); err != nil {
		return "", fmt.Errorf("err create template: %v", err)
	}
	return templateBuf.String(), nil
}

type ServiceConfig struct {
	DomainPascalCase string
	DomainCamelCase  string
	DomainShort      string
	ModuleName       string
}

func NewServiceConfig(domain string) ServiceConfig {
	return ServiceConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		DomainShort:      getDomainShort(domain),
	}
}

var serviceTemplate = `package service

import (
	"context"
	"fmt"

	"{{.ModuleName}}/internal/repository"
)

type {{.DomainPascalCase}} interface {
	// Bar blablabla
	Bar(ctx context.Context) error
}

type {{.DomainCamelCase}}Service struct {
	{{.DomainCamelCase}}Repository repository.{{.DomainPascalCase}}
}

func New{{.DomainPascalCase}}Service({{.DomainCamelCase}}Repository repository.{{.DomainPascalCase}}) {{.DomainPascalCase}} {
	return &{{.DomainCamelCase}}Service{
		{{.DomainCamelCase}}Repository: {{.DomainCamelCase}}Repository,
	}
}

// Bar blablabla.
func ({{.DomainShort}}s *{{.DomainCamelCase}}Service) Bar(ctx context.Context) error {
	if err := {{.DomainShort}}s.{{.DomainCamelCase}}Repository.Foo(ctx); err != nil {
		return fmt.Errorf("err babibu: %v", err)
	}

	if err := {{.DomainShort}}s.{{.DomainCamelCase}}Repository.Baz(ctx); err != nil {
		return fmt.Errorf("err zzzzzz: %v", err)
	}

	return nil
}
`
