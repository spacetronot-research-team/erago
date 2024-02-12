package templ

var Router = `package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(ginEngine *gin.Engine, db *gorm.DB) {
	{{.DomainCamelCase}}Controller := get{{.DomainPascalCase}}Controller(db)

	ginEngine.GET("", {{.DomainCamelCase}}Controller.Qux)
}
`
