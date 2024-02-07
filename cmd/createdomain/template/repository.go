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

func GetRepositoryTemplate(domain string, varErr1 string, varErr2 string) (string, error) {
	repositoryConfig := NewRepositoryConfig(domain, varErr1, varErr2)

	repositoryTemplate, err := template.New("repositoryTemplate").Parse(repositoryTemplate)
	if err != nil {
		return "", fmt.Errorf("err parse template repositoryTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	if err := repositoryTemplate.Execute(&templateBuf, repositoryConfig); err != nil {
		return "", fmt.Errorf("err create template: %v", err)
	}

	return templateBuf.String(), nil
}

type RepositoryConfig struct {
	DomainPascalCase string
	DomainCamelCase  string
	DomainShort      string
	DomainSnakeCase  string
	VarErr1          string
	VarErr2          string
	UniqueErrCode1   string
	UniqueErrCode2   string
}

func NewRepositoryConfig(domain string, varErr1 string, varErr2 string) RepositoryConfig {
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

	return RepositoryConfig{
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

var repositoryTemplate = `package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

//go:generate mockgen -source={{.DomainSnakeCase}}.go -destination=mockrepository/{{.DomainSnakeCase}}.go -package=mockrepository

var (
	Err{{.VarErr1}} = errors.New("[{{.UniqueErrCode1}}] err blabla")
	Err{{.VarErr2}} = errors.New("[{{.UniqueErrCode2}}] err babibu")
)

type {{.DomainPascalCase}} interface {
	// Foo blablabla.
	Foo(ctx context.Context) error
	// Baz blablablabla.
	Baz(ctx context.Context) error
}

type {{.DomainCamelCase}}Repository struct {
	db *gorm.DB
}

func New{{.DomainPascalCase}}Repository(db *gorm.DB) {{.DomainPascalCase}} {
	return &{{.DomainCamelCase}}Repository{
		db: db,
	}
}

// Foo blablablablabla.
func ({{.DomainShort}}r *{{.DomainCamelCase}}Repository) Foo(ctx context.Context) error {
	err := gorm.ErrRecordNotFound // error from query
	if err != nil {
		return errors.Join(err, Err{{.VarErr1}})
	}
	return nil
}

// Baz blablablablabla.
func ({{.DomainShort}}r *{{.DomainCamelCase}}Repository) Baz(ctx context.Context) error {
	err := gorm.ErrRecordNotFound // error from query
	if err != nil {
		return errors.Join(err, Err{{.VarErr2}})
	}
	return nil
}
`
