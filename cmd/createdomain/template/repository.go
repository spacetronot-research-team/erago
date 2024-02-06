package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GetRepositoryTemplate(domain string) (string, error) {
	repositoryConfig := NewRepositoryConfig(domain)

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
}

func NewRepositoryConfig(domain string) RepositoryConfig {
	return RepositoryConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		DomainShort:      getDomainShort(domain),
		DomainSnakeCase:  strcase.ToSnake(domain),
	}
}

var repositoryTemplate = `package repository

import (
	"context"

	"gorm.io/gorm"
)

//go:generate mockgen -source={{.DomainSnakeCase}}.go -destination=mock/{{.DomainSnakeCase}}.go -package=repository

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
	return gorm.ErrRecordNotFound
}

// Baz blablablablabla.
func ({{.DomainShort}}r *{{.DomainCamelCase}}Repository) Baz(ctx context.Context) error {
	return gorm.ErrRecordNotFound
}
`
