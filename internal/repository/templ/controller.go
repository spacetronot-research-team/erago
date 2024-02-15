package templ

//nolint:lll
var Controller = `package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		logrus.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":  "success qux",
		"error": nil,
	})
}
`
