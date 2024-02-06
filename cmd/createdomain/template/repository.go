package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
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
}

func NewRepositoryConfig(domain string, varErr1 string, varErr2 string) RepositoryConfig {
	return RepositoryConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		DomainShort:      getDomainShort(domain),
		DomainSnakeCase:  strcase.ToSnake(domain),
		VarErr1:          varErr1,
		VarErr2:          varErr2,
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
	Err{{.VarErr1}}  = errors.New("err blabla")
	Err{{.VarErr2}} = errors.New("err babibu")
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
