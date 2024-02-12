package templ

var NewInjection = `package router

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
