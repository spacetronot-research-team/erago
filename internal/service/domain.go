package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ettle/strcase"
	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/erago/internal/repository"
	"github.com/spacetronot-research-team/erago/pkg/ctxutil"
	"github.com/spacetronot-research-team/erago/pkg/gomod"
	"github.com/spacetronot-research-team/erago/pkg/osfile"
	"github.com/spacetronot-research-team/erago/pkg/random"
)

//go:generate mockgen -source=domain.go -destination=mockservice/domain.go -package=mockservice

var (
	ErrAddUniqueErrCodeToErrorsJSON = errors.New("err add unique err code to errors.json")
)

type Domain interface {
	// CreateDomain create new domain.
	CreateDomain(ctx context.Context)
}

type domainService struct {
	domainRepository repository.Domain
}

func NewDomainService(domainRepository repository.Domain) Domain {
	return &domainService{
		domainRepository: domainRepository,
	}
}

// CreateDomain implements Domain.
// CreateDomain create new domain.
func (ds *domainService) CreateDomain(ctx context.Context) {
	logrus.Info("create domain start")

	varErr1 := random.StringPascal()
	varErr2 := random.StringPascal()

	ds.generateControllerTemplate(ctx, varErr1)

	ds.generateServiceTemplate(ctx, varErr1, varErr2)

	ds.generateRepositoryTemplate(ctx, varErr1, varErr2)

	ds.generateInjection(ctx)

	isMockgenInstalled := ds.getIsMockgenInstalled(ctx)
	if isMockgenInstalled {
		ds.generateMockRepository(ctx)

		ds.generateMockService(ctx)

		ds.generateRepositoryTestTemplate(ctx, varErr1)

		ds.generateServiceTestTemplate(ctx, varErr1, varErr2)

		ds.generateControllerTestTemplate(ctx, varErr1)
	}

	ds.runGoModTidy()

	if isMockgenInstalled {
		logrus.Info("create domain finish")
	} else {
		logrus.Info("create domain finish without generate mock repository, mock service, and service test")
		logrus.Info("mockgen is not installed, if you want erago to generate mock repository, mock service, and service test please install mockgen.\n\tgo install go.uber.org/mock/mockgen@latest") //nolint:lll
	}

	fileName := strcase.ToSnake(ctxutil.GetDomain(ctx))
	logrus.Info(fmt.Sprintf("controller:\t\t%s:0", filepath.Join(ctxutil.GetControllerDirPath(ctx), fileName+".go")))
	logrus.Info(fmt.Sprintf("service:\t\t%s:0", filepath.Join(ctxutil.GetServiceDirPath(ctx), fileName+".go:0")))
	if isMockgenInstalled {
		logrus.Info(fmt.Sprintf("service test:\t%s:0", filepath.Join(ctxutil.GetServiceDirPath(ctx), fileName+"_test.go")))
	}
	logrus.Info(fmt.Sprintf("repository:\t\t%s:0", filepath.Join(ctxutil.GetRepositoryDirPath(ctx), fileName+".go")))
	logrus.Info(fmt.Sprintf("injection:\t\t%s:0", filepath.Join(ctxutil.GetRouterDirPath(ctx), "injection.go")))
}

func (ds *domainService) runGoModTidy() {
	logrus.Info("run go mod tidy start")
	if err := gomod.RunGoModTidy(); err != nil {
		logrus.Warn(fmt.Errorf("err run go mod tidy: %v", err))
		return
	}
	logrus.Info("run go mod tidy finish")
}

// generateControllerTemplate generate controller template.
func (ds *domainService) generateControllerTemplate(ctx context.Context, varErr1 string) {
	logrus.Info("generate controller template start")

	uniqueErrCode1 := fmt.Sprintf("%s@%s", ctxutil.GetProjectName(ctx), random.String())

	controllerTemplate, err := ds.domainRepository.GetControllerTemplate(ctx, varErr1, uniqueErrCode1)
	if err != nil {
		logrus.Warn(fmt.Errorf("err repo get controller template: %v", err))
		return
	}

	if err := ds.writeControllerTemplateFile(ctx, controllerTemplate); err != nil {
		logrus.Warn(fmt.Errorf("err write controller template file: %v", err))
		return
	}

	if err := osfile.AddUniqueErrCodeToErrorsJSON(ctx, uniqueErrCode1); err != nil {
		logrus.Warn(fmt.Errorf("%v: %v", ErrAddUniqueErrCodeToErrorsJSON, err))
		return
	}

	logrus.Info("generate controller template finish")
}

func (ds *domainService) writeControllerTemplateFile(ctx context.Context, controllerTemplate string) error {
	controllerDirPath := ctxutil.GetControllerDirPath(ctx)
	domainSnakeCase := strcase.ToSnake(ctxutil.GetDomain(ctx))
	path := filepath.Join(controllerDirPath, domainSnakeCase+".go")

	if err := os.WriteFile(path, []byte(controllerTemplate), os.ModePerm); err != nil {
		return fmt.Errorf("err write controller template: %v", err)
	}

	return nil
}

// generateServiceTemplate generate service template.
func (ds *domainService) generateServiceTemplate(ctx context.Context, varErr1 string, varErr2 string) {
	logrus.Info("generate service template start")

	serviceTemplate, err := ds.domainRepository.GetServiceTemplate(ctx, varErr1, varErr2)
	if err != nil {
		logrus.Warn(fmt.Errorf("err repo get service template: %v", err))
		return
	}

	if err := ds.writeServiceTemplateFile(ctx, serviceTemplate); err != nil {
		logrus.Warn(fmt.Errorf("err write service template file: %v", err))
		return
	}

	logrus.Info("generate service template finish")
}

func (ds *domainService) writeServiceTemplateFile(ctx context.Context, serviceTemplate string) error {
	serviceDirPath := ctxutil.GetServiceDirPath(ctx)
	domainSnakeCase := strcase.ToSnake(ctxutil.GetDomain(ctx))
	path := filepath.Join(serviceDirPath, domainSnakeCase+".go")

	if err := os.WriteFile(path, []byte(serviceTemplate), os.ModePerm); err != nil {
		return fmt.Errorf("err write service template: %v", err)
	}

	return nil
}

// generateRepositoryTemplate generate repository template.
func (ds *domainService) generateRepositoryTemplate(ctx context.Context, varErr1 string, varErr2 string) {
	logrus.Info("generate repository template start")

	repositoryTemplate, err := ds.domainRepository.GetRepositoryTemplate(ctx, varErr1, varErr2)
	if err != nil {
		logrus.Warn(fmt.Errorf("err repo get repository template: %v", err))
		return
	}

	if err := ds.writeRepositoryTemplateFile(ctx, repositoryTemplate); err != nil {
		logrus.Warn(fmt.Errorf("err write repository template file: %v", err))
		return
	}

	logrus.Info("generate repository template finish")
}

func (ds *domainService) writeRepositoryTemplateFile(ctx context.Context, repositoryTemplate string) error {
	repositoryDirPath := ctxutil.GetRepositoryDirPath(ctx)
	domainSnakeCase := strcase.ToSnake(ctxutil.GetDomain(ctx))
	path := filepath.Join(repositoryDirPath, domainSnakeCase+".go")

	if err := os.WriteFile(path, []byte(repositoryTemplate), os.ModePerm); err != nil {
		return fmt.Errorf("err write repository template: %v", err)
	}

	return nil
}

// generateInjection generate new injection template or append.
func (ds *domainService) generateInjection(ctx context.Context) {
	logrus.Info("generate injection start")

	injectionFilePath := ctxutil.GetInjectionFilePath(ctx)

	_, err := os.Stat(injectionFilePath)
	if err == nil {
		logrus.Info("injection file found, will append")
		if err := ds.appendInjectionTemplate(ctx); err != nil {
			logrus.Warn(fmt.Errorf("err append injection template: %v", err))
			return
		}
		logrus.Info("generate injection finish")
		return
	}

	if os.IsNotExist(err) {
		logrus.Info("injection file not found, will create new")
		if err := ds.generateNewInjectionTemplate(ctx); err != nil {
			logrus.Warn(fmt.Errorf("err generate new injection template: %v", err))
			return
		}
		logrus.Info("generate injection finish")
		return
	}

	logrus.Warn(fmt.Errorf("err get injection.go file info: %v", err))
}

// appendInjectionTemplate append injection template.
func (ds *domainService) appendInjectionTemplate(ctx context.Context) error {
	injectionFilePath := ctxutil.GetInjectionFilePath(ctx)
	injectionFile, err := os.OpenFile(injectionFilePath, os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("err open injection.go file info: %v", err)
	}
	defer injectionFile.Close()

	appendInjectionTemplate, err := ds.domainRepository.GetAppendInjectionTemplate(ctx)
	if err != nil {
		return fmt.Errorf("err repo get append injection template: %v", err)
	}

	if _, err := injectionFile.WriteString(appendInjectionTemplate); err != nil {
		return fmt.Errorf("err write append injection template: %v", err)
	}

	return nil
}

// generateNewInjectionTemplate generate new injection template.
func (ds *domainService) generateNewInjectionTemplate(ctx context.Context) error {
	newInjectionTemplate, err := ds.domainRepository.GetNewInjectionTemplate(ctx)
	if err != nil {
		return fmt.Errorf("err repo get new injection template: %v", err)
	}

	injectionFilePath := ctxutil.GetInjectionFilePath(ctx)
	if err = os.WriteFile(injectionFilePath, []byte(newInjectionTemplate), os.ModePerm); err != nil {
		return fmt.Errorf("err write new injection template: %v", err)
	}

	return nil
}

func (ds *domainService) getIsMockgenInstalled(ctx context.Context) bool {
	cmd := exec.CommandContext(ctx, "which", "mockgen")

	output, err := cmd.Output()
	if err != nil {
		return false
	}

	mockgenPath := strings.TrimSpace(string(output))
	return mockgenPath != ""
}

// generateMockRepository generate mock repository using mockgen.
func (ds *domainService) generateMockRepository(ctx context.Context) {
	logrus.Info("generate mock repository using mockgen start")

	domain := ctxutil.GetDomain(ctx)

	domainFileName := strcase.ToSnake(domain)

	source := fmt.Sprintf("-source=%s.go", domainFileName)
	destination := fmt.Sprintf("-destination=%s", filepath.Join("mock", domainFileName+".go"))

	cmd := exec.Command("mockgen", source, destination, "-package=mock")
	cmd.Dir = ctxutil.GetRepositoryDirPath(ctx)
	stdout, err := cmd.Output()
	if err != nil {
		logrus.Warn(fmt.Errorf("err run mockgen repository: %v", err))
		return
	}

	if string(stdout) != "" {
		fmt.Println(string(stdout)) //nolint:forbidigo
	}

	logrus.Info("generate mock repository using mockgen finish")
}

// generateMockService generate mock repository using mockgen.
func (ds *domainService) generateMockService(ctx context.Context) {
	logrus.Info("generate mock service using mockgen start")

	domain := ctxutil.GetDomain(ctx)

	domainFileName := strcase.ToSnake(domain)

	source := fmt.Sprintf("-source=%s.go", domainFileName)
	destination := fmt.Sprintf("-destination=%s", filepath.Join("mock", domainFileName+".go"))

	cmd := exec.Command("mockgen", source, destination, "-package=mock")
	cmd.Dir = ctxutil.GetServiceDirPath(ctx)
	stdout, err := cmd.Output()
	if err != nil {
		logrus.Warn(fmt.Errorf("err run mockgen service: %v", err))
		return
	}

	if string(stdout) != "" {
		fmt.Println(string(stdout)) //nolint:forbidigo
	}

	logrus.Info("generate mock service using mockgen finish")
}

// generateServiceTestTemplate generate service test template.
func (ds *domainService) generateServiceTestTemplate(ctx context.Context, varErr1 string, varErr2 string) {
	logrus.Info("generate service test template start")

	serviceTestTemplate, err := ds.domainRepository.GetServiceTestTemplate(ctx, varErr1, varErr2)
	if err != nil {
		logrus.Warn(fmt.Errorf("err repo get service test template: %v", err))
		return
	}

	if err := ds.writeServiceTestTemplateFile(ctx, serviceTestTemplate); err != nil {
		logrus.Warn(fmt.Errorf("err write service test template file: %v", err))
		return
	}

	logrus.Info("generate service test template finish")
}

func (ds *domainService) writeServiceTestTemplateFile(ctx context.Context, serviceTestTemplate string) error {
	serviceDirPath := ctxutil.GetServiceDirPath(ctx)
	domainSnakeCase := strcase.ToSnake(ctxutil.GetDomain(ctx))
	path := filepath.Join(serviceDirPath, domainSnakeCase+"_test.go")

	if err := os.WriteFile(path, []byte(serviceTestTemplate), os.ModePerm); err != nil {
		return fmt.Errorf("err write service test template: %v", err)
	}

	return nil
}

// generateControllerTestTemplate generate controller test template.
func (ds *domainService) generateControllerTestTemplate(ctx context.Context, varErr1 string) {
	logrus.Info("generate controller test template start")

	controllerTestTemplate, err := ds.domainRepository.GetControllerTestTemplate(ctx, varErr1)
	if err != nil {
		logrus.Warn(fmt.Errorf("err repo get controller test template: %v", err))
		return
	}

	if err := ds.writeControllerTestTemplateFile(ctx, controllerTestTemplate); err != nil {
		logrus.Warn(fmt.Errorf("err write controller test template file: %v", err))
		return
	}

	logrus.Info("generate controller test template finish")
}

func (ds *domainService) writeControllerTestTemplateFile(ctx context.Context, controllerTestTemplate string) error {
	controllerDirPath := ctxutil.GetControllerDirPath(ctx)
	domainSnakeCase := strcase.ToSnake(ctxutil.GetDomain(ctx))
	path := filepath.Join(controllerDirPath, domainSnakeCase+"_test.go")

	if err := os.WriteFile(path, []byte(controllerTestTemplate), os.ModePerm); err != nil {
		return fmt.Errorf("err write controller test template: %v", err)
	}

	return nil
}

// generateRepositoryTestTemplate generate repository test template.
func (ds *domainService) generateRepositoryTestTemplate(ctx context.Context, varErr1 string) {
	logrus.Info("generate repository test template start")

	repositoryTestTemplate, err := ds.domainRepository.GetRepositoryTestTemplate(ctx, varErr1)
	if err != nil {
		logrus.Warn(fmt.Errorf("err repo get repository test template: %v", err))
		return
	}

	if err := ds.writeRepositoryTestTemplateFile(ctx, repositoryTestTemplate); err != nil {
		logrus.Warn(fmt.Errorf("err write repository test template file: %v", err))
		return
	}

	logrus.Info("generate repository test template finish")
}

func (ds *domainService) writeRepositoryTestTemplateFile(ctx context.Context, repositoryTestTemplate string) error {
	repositoryDirPath := ctxutil.GetRepositoryDirPath(ctx)
	domainSnakeCase := strcase.ToSnake(ctxutil.GetDomain(ctx))
	path := filepath.Join(repositoryDirPath, domainSnakeCase+"_test.go")

	if err := os.WriteFile(path, []byte(repositoryTestTemplate), os.ModePerm); err != nil {
		return fmt.Errorf("err write repository test template: %v", err)
	}

	return nil
}
