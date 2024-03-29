package templ

var Repository = `package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

//go:generate mockgen -source={{.DomainSnakeCase}}.go -destination=mock/{{.DomainSnakeCase}}.go -package=mock

var (
	Err{{.VarErr1}} = errors.New("err blabla")
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
	var quuz int32
	err := {{.DomainShort}}r.db.Raw("SELECT 1").Scan(&quuz).Error
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
