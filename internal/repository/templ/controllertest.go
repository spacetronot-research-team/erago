package templ

//nolint:dupword
var ControllerTest = `package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"{{.ModuleName}}/internal/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test{{.DomainPascalCase}}Controller_Qux(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type fields struct {
		{{.DomainCamelCase}}Service *mock.Mock{{.DomainPascalCase}}
	}
	type args struct {
		reqHeader map[string]string
		reqBody gin.H
	}
	tests := []struct {
		name     string
		mock     func(f fields)
		args     args
		wantCode int
		wantBody gin.H
	}{
		{
			name: "Success",
			mock: func(f fields) {
				f.{{.DomainCamelCase}}Service.EXPECT().
					Bar(gomock.Any()).
					Return(nil)
			},
			args:     args{},
			wantCode: http.StatusOK,
			wantBody: gin.H{
				"data":  "success qux",
				"error": nil,
			},
		},
		{
			name: "Bad Request",
			mock: func(f fields) {
				f.{{.DomainCamelCase}}Service.EXPECT().
					Bar(gomock.Any()).
					Return(assert.AnError)
			},
			args:     args{},
			wantCode: http.StatusBadRequest,
			wantBody: gin.H{
				"data":  nil,
				"error": errors.Join(assert.AnError, Err{{.VarErr1}}).Error(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				{{.DomainCamelCase}}Service: mock.NewMock{{.DomainPascalCase}}(ctrl),
			}
			tt.mock(f)

			{{.DomainShort}} := &{{.DomainPascalCase}}Controller{
				{{.DomainCamelCase}}Service: f.{{.DomainCamelCase}}Service,
			}

			rr := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rr)
			reqBody, _ := json.Marshal(tt.args.reqBody)
			req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(reqBody))
			for k, v := range tt.args.reqHeader {
				req.Header.Set(k, v)
			}
			ctx.Request = req

			{{.DomainShort}}.Qux(ctx)

			assert.Equal(t, tt.wantCode, rr.Code)

			wantBody, _ := json.Marshal(tt.wantBody)
			assert.Equal(t, string(wantBody), rr.Body.String())
		})
	}
}
`
