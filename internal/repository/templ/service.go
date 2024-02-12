package templ

//nolint:lll
var Service = `package service

import (
	"context"
	"errors"

	"{{.ModuleName}}/internal/repository"
)

//go:generate mockgen -source={{.DomainSnakeCase}}.go -destination=mock/{{.DomainSnakeCase}}.go -package=mock

var (
	Err{{.VarErr1}} = errors.New("err jasdfsefs")
	Err{{.VarErr2}} = errors.New("err jasdf")
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
