package template

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/erago/common/gomod"
	"github.com/spacetronot-research-team/erago/common/osfile"
	"github.com/spacetronot-research-team/erago/common/random"
)

func GetServiceTemplate(domain string, varErr1 string, varErr2 string) (string, error) {
	serviceConfig := NewServiceConfig(domain, varErr1, varErr2)
	moduleName, err := gomod.GetModuleName()
	if err != nil {
		return "", fmt.Errorf("err get module name: %v", err)
	}
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
	DomainSnakeCase  string
	ModuleName       string
	VarErr1          string
	VarErr2          string
	UniqueErrCode1   string
	UniqueErrCode2   string
}

func NewServiceConfig(domain string, varErr1 string, varErr2 string) ServiceConfig {
	projectName, err := gomod.GetProjectNameFromModule()
	if err != nil {
		logrus.Warn(fmt.Errorf("err get project name from module: %v", err))
	}

	uniqueErrCode1 := fmt.Sprintf("%s@%s", projectName, random.String())
	uniqueErrCode2 := fmt.Sprintf("%s@%s", projectName, random.String())
	jsonFilePath := filepath.Join("docs", "errors.json")
	err = osfile.AddUniqueErrCodeToErrorsJSON(jsonFilePath, uniqueErrCode1, uniqueErrCode2)
	if err != nil {
		logrus.Warn(fmt.Errorf("err add unique err code to errors.json: %v", err))
	}

	return ServiceConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		DomainShort:      getDomainShort(domain),
		DomainSnakeCase:  strcase.ToSnake(domain),
		VarErr1:          varErr1,
		VarErr2:          varErr2,
		UniqueErrCode1:   uniqueErrCode1,
		UniqueErrCode2:   uniqueErrCode2,
	}
}

var serviceTemplate = `package service

import (
	"context"
	"errors"

	"{{.ModuleName}}/internal/repository"
)

//go:generate mockgen -source={{.DomainSnakeCase}}.go -destination=mockservice/{{.DomainSnakeCase}}.go -package=mockservice

var (
	Err{{.VarErr1}} = errors.New("[{{.UniqueErrCode1}}] err jasdfsefs")
	Err{{.VarErr2}} = errors.New("[{{.UniqueErrCode2}}] err jasdf")
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
