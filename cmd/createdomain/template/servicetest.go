package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GteServiceTestTemplate(domain string, moduleName string, varErr1 string, varErr2 string) (string, error) {
	serviceTestConfig := NewServiceTestConfig(domain, varErr1, varErr2)
	serviceTestConfig.ModuleName = moduleName

	serviceTestTemplate, err := template.New("serviceTestTemplate").Parse(serviceTestTemplate)
	if err != nil {
		return "", fmt.Errorf("err parse template serviceTestTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	if err = serviceTestTemplate.Execute(&templateBuf, serviceTestConfig); err != nil {
		return "", fmt.Errorf("err create template: %v", err)
	}

	return templateBuf.String(), nil
}

type ServiceTestConfig struct {
	DomainPascalCase string
	DomainCamelCase  string
	DomainShort      string
	ModuleName       string
	VarErr1          string
	VarErr2          string
}

func NewServiceTestConfig(domain string, varErr1 string, varErr2 string) ServiceTestConfig {
	return ServiceTestConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		DomainShort:      getDomainShort(domain),
		VarErr1:          varErr1,
		VarErr2:          varErr2,
	}
}

var serviceTestTemplate = `package service

import (
	"context"
	"testing"

	"{{.ModuleName}}/internal/repository/mockrepository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_{{.DomainCamelCase}}Service_Bar(t *testing.T) {
	type fields struct {
		{{.DomainCamelCase}}Repository *mockrepository.Mock{{.DomainPascalCase}}
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		mock    func(f fields)
		args    args
		wantErr error
	}{
		{
			name: "bar err foo",
			mock: func(f fields) {
				f.{{.DomainCamelCase}}Repository.EXPECT().
					Foo(nil).Return(assert.AnError)
			},
			args: args{
				ctx: nil,
			},
			wantErr: Err{{.VarErr1}},
		},
		{
			name: "bar err baz",
			mock: func(f fields) {
				f.{{.DomainCamelCase}}Repository.EXPECT().
					Foo(nil).Return(nil)

				f.{{.DomainCamelCase}}Repository.EXPECT().
					Baz(nil).Return(assert.AnError)
			},
			args: args{
				ctx: nil,
			},
			wantErr: Err{{.VarErr2}},
		},
		{
			name: "bar success",
			mock: func(f fields) {
				f.{{.DomainCamelCase}}Repository.EXPECT().
					Foo(nil).Return(nil)

				f.{{.DomainCamelCase}}Repository.EXPECT().
					Baz(nil).Return(nil)
			},
			args: args{
				ctx: nil,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				{{.DomainCamelCase}}Repository: mockrepository.NewMock{{.DomainPascalCase}}(ctrl),
			}
			tt.mock(f)

			{{.DomainShort}}s := &{{.DomainCamelCase}}Service{
				{{.DomainCamelCase}}Repository: f.{{.DomainCamelCase}}Repository,
			}

			err := {{.DomainShort}}s.Bar(tt.args.ctx)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
`
