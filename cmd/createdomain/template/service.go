package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GetServiceTemplate(domain string, moduleName string, varErr1 string, varErr2 string) (string, error) {
	serviceConfig := NewServiceConfig(domain, varErr1, varErr2)
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
	VarErr1          string
	VarErr2          string
}

func NewServiceConfig(domain string, varErr1 string, varErr2 string) ServiceConfig {
	return ServiceConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		DomainShort:      getDomainShort(domain),
		VarErr1:          varErr1,
		VarErr2:          varErr2,
	}
}

var serviceTemplate = `package service

import (
	"context"
	"errors"

	"{{.ModuleName}}/internal/repository"
)

var (
	Err{{.VarErr1}} = errors.New("err jasdfsefs")
	Err{{.VarErr2}}  = errors.New("err jasdf")
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
		return errors.Join(err, Err{{.VarErr1}})
	}

	if err := {{.DomainShort}}s.{{.DomainCamelCase}}Repository.Baz(ctx); err != nil {
		return errors.Join(err, Err{{.VarErr2}})
	}

	return nil
}
`
