package templ

//nolint:lll
var Controller = `package http

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
