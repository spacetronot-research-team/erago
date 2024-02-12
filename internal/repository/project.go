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

//go:generate mockgen -source=project.go -destination=mockrepository/project.go -package=mockrepository

type Project interface {
	// GetCMDMainTemplate return string cmd main template.
	GetCMDMainTemplate(ctx context.Context) (string, error)
	// GetRouterTemplate return string router template.
	GetRouterTemplate(ctx context.Context) (string, error)
	// GetReadmeTemplate return string router template.
	GetReadmeTemplate(ctx context.Context) (string, error)
}

type projectRepository struct {
}

func NewProjectRepository() Project {
	return &projectRepository{}
}

// GetCMDMainTemplate return string cmd main template.
func (*projectRepository) GetCMDMainTemplate(ctx context.Context) (string, error) {
	cmdMainTemplate, err := template.New("cmdMainTemplate").Parse(templ.CMDMain)
	if err != nil {
		return "", fmt.Errorf("err parse template cmdMainTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	cmdMainConfig := NewCMDMainConfig(ctx)
	if err = cmdMainTemplate.Execute(&templateBuf, cmdMainConfig); err != nil {
		return "", fmt.Errorf("err execute cmd main template: %v", err)
	}

	return templateBuf.String(), nil
}

type CMDMainConfig struct {
	ModuleName string
}

func NewCMDMainConfig(ctx context.Context) CMDMainConfig {
	return CMDMainConfig{
		ModuleName: ctxutil.GetModuleName(ctx),
	}
}

// GetRouterTemplate return string cmd main template.
func (*projectRepository) GetRouterTemplate(ctx context.Context) (string, error) {
	routerTemplate, err := template.New("routerTemplate").Parse(templ.Router)
	if err != nil {
		return "", fmt.Errorf("err parse template routerTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	routerConfig := NewRouterConfig(ctx)
	if err = routerTemplate.Execute(&templateBuf, routerConfig); err != nil {
		return "", fmt.Errorf("err execute router template: %v", err)
	}

	return templateBuf.String(), nil
}

type RouterConfig struct {
	DomainPascalCase string
	DomainCamelCase  string
}

func NewRouterConfig(ctx context.Context) RouterConfig {
	domain := ctxutil.GetDomain(ctx)
	return RouterConfig{
		DomainPascalCase: strcase.ToCamel(domain),
		DomainCamelCase:  strcase.ToLowerCamel(domain),
	}
}

// GetReadmeTemplate return string readme template.
func (*projectRepository) GetReadmeTemplate(ctx context.Context) (string, error) {
	readmeTemplate, err := template.New("readmeTemplate").Parse(templ.GetReadmeTemplate())
	if err != nil {
		return "", fmt.Errorf("err parse template readmeTemplate: %v", err)
	}

	var templateBuf bytes.Buffer
	readmeConfig := NewReadmeConfig(ctx)
	if err = readmeTemplate.Execute(&templateBuf, readmeConfig); err != nil {
		return "", fmt.Errorf("err execute readme template: %v", err)
	}

	return templateBuf.String(), nil
}

type ReadmeConfig struct {
	ProjectName string
}

func NewReadmeConfig(ctx context.Context) ReadmeConfig {
	return ReadmeConfig{
		ProjectName: ctxutil.GetProjectName(ctx),
	}
}
