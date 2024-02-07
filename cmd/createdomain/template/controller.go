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

func GetControllerTemplate(domain string, varErr1 string) (string, error) {
	controllerConfig := NewControllerConfig(domain, varErr1)
	moduleName, err := gomod.GetModuleName()
	if err != nil {
		return "", fmt.Errorf("err get module name: %v", err)
	}
	controllerConfig.ModuleName = moduleName

	controllerTemplate, err := template.New("controllerTemplate").Parse(controllerTemplate)
	if err != nil {
		return "", fmt.Errorf("err parse template controllerTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	if err = controllerTemplate.Execute(&templateBuf, controllerConfig); err != nil {
		return "", fmt.Errorf("err create template: %v", err)
	}

	return templateBuf.String(), nil
}

type ControllerConfig struct {
	DomainPascalCase string
	DomainCamelCase  string
	DomainShort      string
	ModuleName       string
	VarErr1          string
	UniqueErrCode1   string
}

func NewControllerConfig(domain string, varErr1 string) ControllerConfig {
	projectName, err := gomod.GetProjectNameFromModule()
	if err != nil {
		logrus.Warn(fmt.Errorf("err get project name from module: %v", err))
	}

	uniqueErrCode1 := fmt.Sprintf("%s@%s", projectName, random.String())
	jsonFilePath := filepath.Join("docs", "errors.json")
	err = osfile.AddUniqueErrCodeToErrorsJSON(jsonFilePath, uniqueErrCode1)
	if err != nil {
		logrus.Warn(fmt.Errorf("err add unique err code to errors.json: %v", err))
	}

	return ControllerConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		DomainShort:      getDomainShort(domain),
		VarErr1:          varErr1,
		UniqueErrCode1:   uniqueErrCode1,
	}
}

var controllerTemplate = `package http

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"{{.ModuleName}}/internal/service"
)

var (
	Err{{.VarErr1}} = errors.New("[{{.UniqueErrCode1}}] err jklasjd")
)

type {{.DomainPascalCase}}Controller struct {
	{{.DomainCamelCase}}Service service.{{.DomainPascalCase}}
}

func New{{.DomainPascalCase}}Controller({{.DomainCamelCase}}Service service.{{.DomainPascalCase}}) *{{.DomainPascalCase}}Controller {
	return &{{.DomainPascalCase}}Controller{
		{{.DomainCamelCase}}Service: {{.DomainCamelCase}}Service,
	}
}

// Qux babibu.
func ({{.DomainShort}}c *{{.DomainPascalCase}}Controller) Qux(ctx *gin.Context) {
	if err := {{.DomainShort}}c.{{.DomainCamelCase}}Service.Bar(ctx); err != nil {
		err = errors.Join(err, Err{{.VarErr1}})
		log.Println(err)
		return
	}
	log.Println("^.^")
}
`
