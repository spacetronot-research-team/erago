package repository

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/spacetronot-research-team/erago/internal/repository/templ"
	"github.com/spacetronot-research-team/erago/pkg/ctxutil"
)

//go:generate mockgen -source=domain.go -destination=mockrepository/domain.go -package=mockrepository

type Domain interface {
	// GetControllerTemplate return string controller template.
	GetControllerTemplate(ctx context.Context, varErr1 string, uniqueErrCode1 string) (string, error)
	// GetControllerTestTemplate return string controller template.
	GetControllerTestTemplate(ctx context.Context, varErr1 string) (string, error)
	// GetServiceTemplate return string service template.
	GetServiceTemplate(ctx context.Context, varErr1 string, varErr2 string) (string, error)
	// GetServiceTestTemplate return string service test template.
	GetServiceTestTemplate(ctx context.Context, varErr1 string, varErr2 string) (string, error)
	// GetRepositoryTemplate return string repository template.
	GetRepositoryTemplate(ctx context.Context, varErr1 string, varErr2 string) (string, error)
	// GetRepositoryTestTemplate return string repository test template.
	GetRepositoryTestTemplate(ctx context.Context, varErr1 string) (string, error)
	// GetNewInjectionTemplate return string new injection template.
	GetNewInjectionTemplate(ctx context.Context) (string, error)
	// GetAppendInjectionTemplate return string append injection template.
	GetAppendInjectionTemplate(ctx context.Context) (string, error)
}

type domainRepository struct {
}

func NewDomainRepository() Domain {
	return &domainRepository{}
}

// GetControllerTemplate implements Domain.
// GetControllerTemplate return string controller template.
func (*domainRepository) GetControllerTemplate(ctx context.Context, varErr1 string, uniqueErrCode1 string) (string, error) { //nolint:lll
	controllerTemplate, err := template.New("controllerTemplate").Parse(templ.Controller)
	if err != nil {
		return "", fmt.Errorf("err parse template controllerTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	controllerConfig := NewControllerConfig(ctx, varErr1, uniqueErrCode1)
	if err = controllerTemplate.Execute(&templateBuf, controllerConfig); err != nil {
		return "", fmt.Errorf("err execute controller template: %v", err)
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

func NewControllerConfig(ctx context.Context, varErr1 string, uniqueErrCode1 string) ControllerConfig {
	domain := ctxutil.GetDomain(ctx)
	domainShort := ctxutil.GetDomainShort(ctx)
	moduleName := ctxutil.GetModuleName(ctx)

	return ControllerConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		DomainShort:      domainShort,
		ModuleName:       moduleName,
		VarErr1:          varErr1,
		UniqueErrCode1:   uniqueErrCode1,
	}
}

// GetControllerTestTemplate implements Domain.
// GetControllerTestTemplate return string controller test template.
func (*domainRepository) GetControllerTestTemplate(ctx context.Context, varErr1 string) (string, error) {
	controllerTestTemplate, err := template.New("controllerTestTemplate").Parse(templ.ControllerTest)
	if err != nil {
		return "", fmt.Errorf("err parse template ControllerTestTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	controllerTestConfig := NewControllerTestConfig(ctx, varErr1)
	if err = controllerTestTemplate.Execute(&templateBuf, controllerTestConfig); err != nil {
		return "", fmt.Errorf("err execute controller test template: %v", err)
	}

	return templateBuf.String(), nil
}

type ControllerTestConfig struct {
	DomainPascalCase string
	DomainCamelCase  string
	DomainShort      string
	ModuleName       string
	VarErr1          string
}

func NewControllerTestConfig(ctx context.Context, varErr1 string) ControllerTestConfig {
	domain := ctxutil.GetDomain(ctx)
	domainShort := ctxutil.GetDomainShort(ctx)
	moduleName := ctxutil.GetModuleName(ctx)

	return ControllerTestConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		DomainShort:      domainShort,
		ModuleName:       moduleName,
		VarErr1:          varErr1,
	}
}

// GetServiceTemplate implements Domain.
// GetServiceTemplate return string service template.
func (*domainRepository) GetServiceTemplate(ctx context.Context, varErr1 string, varErr2 string) (string, error) { //nolint:lll
	serviceTemplate, err := template.New("serviceTemplate").Parse(templ.Service)
	if err != nil {
		return "", fmt.Errorf("err parse template serviceTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	serviceConfig := NewServiceConfig(ctx, varErr1, varErr2)
	if err = serviceTemplate.Execute(&templateBuf, serviceConfig); err != nil {
		return "", fmt.Errorf("err execute service template: %v", err)
	}

	return templateBuf.String(), nil
}

type ServiceConfig struct {
	DomainPascalCase string
	DomainCamelCase  string
	DomainShort      string
	DomainSnakeCase  string
	ModuleName       string
	VarErr1          string
	VarErr2          string
}

func NewServiceConfig(ctx context.Context, varErr1 string, varErr2 string) ServiceConfig {
	domain := ctxutil.GetDomain(ctx)
	domainShort := ctxutil.GetDomainShort(ctx)
	moduleName := ctxutil.GetModuleName(ctx)

	return ServiceConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		DomainShort:      domainShort,
		DomainSnakeCase:  strcase.ToSnake(domain),
		ModuleName:       moduleName,
		VarErr1:          varErr1,
		VarErr2:          varErr2,
	}
}

// GetRepositoryTemplate implements Domain.
// GetRepositoryTemplate return string repository template.
func (*domainRepository) GetRepositoryTemplate(ctx context.Context, varErr1 string, varErr2 string) (string, error) { //nolint:lll
	repositoryTemplate, err := template.New("RepositoryTemplate").Parse(templ.Repository)
	if err != nil {
		return "", fmt.Errorf("err parse template repositoryTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	repositoryConfig := NewRepositoryConfig(ctx, varErr1, varErr2)
	if err = repositoryTemplate.Execute(&templateBuf, repositoryConfig); err != nil {
		return "", fmt.Errorf("err execute repository template: %v", err)
	}

	return templateBuf.String(), nil
}

type RepositoryConfig struct {
	DomainPascalCase string
	DomainCamelCase  string
	DomainShort      string
	DomainSnakeCase  string
	VarErr1          string
	VarErr2          string
}

func NewRepositoryConfig(ctx context.Context, varErr1 string, varErr2 string) RepositoryConfig {
	domain := ctxutil.GetDomain(ctx)
	domainShort := ctxutil.GetDomainShort(ctx)

	return RepositoryConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		DomainShort:      domainShort,
		DomainSnakeCase:  strcase.ToSnake(domain),
		VarErr1:          varErr1,
		VarErr2:          varErr2,
	}
}

// GetNewInjectionTemplate implements Domain.
// GetNewInjectionTemplate return string new injection template.
func (*domainRepository) GetNewInjectionTemplate(ctx context.Context) (string, error) {
	newInjectionTemplate, err := template.New("newInjectionTemplate").Parse(templ.NewInjection)
	if err != nil {
		return "", fmt.Errorf("err parse template newInjectionTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	newInjectionConfig := NewNewInjectionConfig(ctx)
	if err = newInjectionTemplate.Execute(&templateBuf, newInjectionConfig); err != nil {
		return "", fmt.Errorf("err execute new injection template: %v", err)
	}

	return templateBuf.String(), nil
}

type NewInjectionConfig struct {
	DomainPascalCase string
	DomainCamelCase  string
	ModuleName       string
}

func NewNewInjectionConfig(ctx context.Context) NewInjectionConfig {
	domain := ctxutil.GetDomain(ctx)
	moduleName := ctxutil.GetModuleName(ctx)

	return NewInjectionConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		ModuleName:       moduleName,
	}
}

// GetAppendInjectionTemplate implements Domain.
// GetAppendInjectionTemplate return string append injection template.
func (*domainRepository) GetAppendInjectionTemplate(ctx context.Context) (string, error) {
	appendInjectionTemplate, err := template.New("AppendInjectionTemplate").Parse(templ.AppendInjection)
	if err != nil {
		return "", fmt.Errorf("err parse template appendInjectionTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	appendInjectionConfig := NewAppendInjectionConfig(ctx)
	if err = appendInjectionTemplate.Execute(&templateBuf, appendInjectionConfig); err != nil {
		return "", fmt.Errorf("err execute append injection template: %v", err)
	}

	return templateBuf.String(), nil
}

type AppendInjectionConfig struct {
	DomainPascalCase string
	DomainCamelCase  string
}

func NewAppendInjectionConfig(ctx context.Context) AppendInjectionConfig {
	domain := ctxutil.GetDomain(ctx)

	return AppendInjectionConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
	}
}

// GetServiceTestTemplate implements Domain.
// GetServiceTestTemplate return string service test template.
func (*domainRepository) GetServiceTestTemplate(ctx context.Context, varErr1 string, varErr2 string) (string, error) {
	serviceTestTemplate, err := template.New("ServiceTestTemplate").Parse(templ.ServiceTest)
	if err != nil {
		return "", fmt.Errorf("err parse template serviceTestTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	serviceTestConfig := NewServiceTestConfig(ctx, varErr1, varErr2)
	if err = serviceTestTemplate.Execute(&templateBuf, serviceTestConfig); err != nil {
		return "", fmt.Errorf("err execute service test template: %v", err)
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

func NewServiceTestConfig(ctx context.Context, varErr1 string, varErr2 string) ServiceTestConfig {
	domain := ctxutil.GetDomain(ctx)
	domainShort := ctxutil.GetDomainShort(ctx)
	moduleName := ctxutil.GetModuleName(ctx)

	return ServiceTestConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
		DomainShort:      domainShort,
		ModuleName:       moduleName,
		VarErr1:          varErr1,
		VarErr2:          varErr2,
	}
}

// GetRepositoryTestTemplate implements Domain.
// GetRepositoryTemplate return string repository test template.
func (*domainRepository) GetRepositoryTestTemplate(ctx context.Context, varErr1 string) (string, error) {
	repositoryTestTemplate, err := template.New("repositoyrTestTemplate").Parse(templ.RepositoryTest)
	if err != nil {
		return "", fmt.Errorf("err parse template repositoryTestTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	repositoryTestConfig := NewRepositoryTestConfig(ctx, varErr1)
	if err := repositoryTestTemplate.Execute(&templateBuf, repositoryTestConfig); err != nil {
		return "", fmt.Errorf("err execute repository template test: %v", err)
	}

	return templateBuf.String(), nil
}

type RepositoryTestConfig struct {
	DomainCamelCase string
	DomainShort     string
	VarErr1         string
}

func NewRepositoryTestConfig(ctx context.Context, varErr1 string) RepositoryTestConfig {
	domain := ctxutil.GetDomain(ctx)
	domainShort := ctxutil.GetDomainShort(ctx)

	return RepositoryTestConfig{
		DomainCamelCase: strcase.ToLowerCamel(domain),
		DomainShort:     domainShort,
		VarErr1:         varErr1,
	}
}
