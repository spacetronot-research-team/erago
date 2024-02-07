package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GetInjectionAppendTemplate(domain string) (string, error) {
	injectionAppendConfig := NewInjectionConfigAppend(domain)

	injectionAppendTemplate, err := template.New("injectionAppendTemplate").Parse(injectionAppendTemplate)
	if err != nil {
		return "", fmt.Errorf("err parse template injectionAppendTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	if err = injectionAppendTemplate.Execute(&templateBuf, injectionAppendConfig); err != nil {
		return "", fmt.Errorf("err create template: %v", err)
	}

	return templateBuf.String(), nil
}

type InjectionConfigAppend struct {
	DomainPascalCase string
	DomainCamelCase  string
}

func NewInjectionConfigAppend(domain string) InjectionConfigAppend {
	return InjectionConfigAppend{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
	}
}

var injectionAppendTemplate = `
func get{{.DomainPascalCase}}Controller(db *gorm.DB) *http.{{.DomainPascalCase}}Controller {
	{{.DomainCamelCase}}Repository := repository.New{{.DomainPascalCase}}Repository(db)
	{{.DomainCamelCase}}Service := service.New{{.DomainPascalCase}}Service({{.DomainCamelCase}}Repository)
	{{.DomainCamelCase}}Controller := http.New{{.DomainPascalCase}}Controller({{.DomainCamelCase}}Service)
	return {{.DomainCamelCase}}Controller
}
`
