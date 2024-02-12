package templ

var AppendInjection = `
func get{{.DomainPascalCase}}Controller(db *gorm.DB) *http.{{.DomainPascalCase}}Controller {
	{{.DomainCamelCase}}Repository := repository.New{{.DomainPascalCase}}Repository(db)
	{{.DomainCamelCase}}Service := service.New{{.DomainPascalCase}}Service({{.DomainCamelCase}}Repository)
	{{.DomainCamelCase}}Controller := http.New{{.DomainPascalCase}}Controller({{.DomainCamelCase}}Service)
	return {{.DomainCamelCase}}Controller
}
`
