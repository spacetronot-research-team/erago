package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
)

func GteServiceTestTemplate(domain string, moduleName string) (string, error) {
	serviceTestConfig := NewServiceTestConfig(domain)
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
}

func NewServiceTestConfig(domain string) ServiceTestConfig {
	return ServiceTestConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		DomainShort:      getDomainShort(domain),
	}
}

var serviceTestTemplate = `package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	repository "{{.ModuleName}}/internal/repository/mock"
)

func Test_{{.DomainCamelCase}}Service_Bar(t *testing.T) {
	type fields struct {
		{{.DomainCamelCase}}Repository *repository.Mock{{.DomainPascalCase}}
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		mock    func(f fields)
		args    args
		wantErr bool
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
			wantErr: true,
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
			wantErr: true,
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				{{.DomainCamelCase}}Repository: repository.NewMock{{.DomainPascalCase}}(ctrl),
			}
			tt.mock(f)

			{{.DomainShort}}s := &{{.DomainCamelCase}}Service{
				{{.DomainCamelCase}}Repository: f.{{.DomainCamelCase}}Repository,
			}

			err := {{.DomainShort}}s.Bar(tt.args.ctx)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
`
