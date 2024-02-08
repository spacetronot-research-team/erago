package createdomain

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/iancoleman/strcase"
	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/erago/cmd/createdomain/template"
	"github.com/spacetronot-research-team/erago/common/gomod"
	"github.com/spacetronot-research-team/erago/common/random"
)

// CreateDomain is main func to create new domain.
func CreateDomain(domain string) {
	logrus.Info("create domain start")

	varErr1 := random.StringPascal()
	varErr2 := random.StringPascal()

	logrus.Info("generate controller template")
	if err := generateControllerTemplate(domain, varErr1); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("generate service template")
	if err := generateServiceTemplate(domain, varErr1, varErr2); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("generate repository template")
	if err := generateRepositoryTemplate(domain, varErr1, varErr2); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("generate injection")
	if err := generateInjectionTemplate(domain); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("generate mock repository using mockgen")
	if err := generateMockRepository(domain); err != nil {
		logrus.Warn("err generate mock repository using mockgen, did you have `mockgen`?")
		logrus.Warn("try to install mockgen:\n\tgo install go.uber.org/mock/mockgen@latest")
		logrus.Warn("create domain finish without mock and test template")

		fileName := strcase.ToSnake(domain)
		logrus.Info(fmt.Sprintf("domain %s controller created:\n\t%s:0", domain, filepath.Join("internal", "controller", "http", fileName+".go")))
		logrus.Info(fmt.Sprintf("domain %s service created:\n\t%s:0", domain, filepath.Join("internal", "service", fileName+".go")))
		logrus.Info(fmt.Sprintf("domain %s repository created:\n\t%s:0", domain, filepath.Join("internal", "repository", fileName+".go")))
		logrus.Info(fmt.Sprintf("domain %s injection created:\n\t%s:0", domain, filepath.Join("internal", "router", "injection.go")))

		return
	}

	logrus.Info("generate service test template")
	if err := generateServiceTestTemplate(domain, varErr1, varErr2); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("generate mock service using mockgen")
	if err := generateMockService(domain); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("run go mod tidy")
	if err := gomod.RunGoModTidy(); err != nil {
		logrus.Fatal(fmt.Errorf("err run go mod tidy: %v", err))
	}

	logrus.Info("create domain finish")

	fileName := strcase.ToSnake(domain)
	logrus.Info(fmt.Sprintf("domain %s controller created:\n\t%s:0", domain, filepath.Join("internal", "controller", "http", fileName+".go")))
	logrus.Info(fmt.Sprintf("domain %s service created:\n\t%s:0", domain, filepath.Join("internal", "service", fileName+".go:0")))
	logrus.Info(fmt.Sprintf("domain %s service test created:\n\t%s:0", domain, filepath.Join("internal", "service", fileName+"_test.go")))
	logrus.Info(fmt.Sprintf("domain %s repository created:\n\t%s:0", domain, filepath.Join("internal", "repository", fileName+".go")))
	logrus.Info(fmt.Sprintf("domain %s injection created:\n\t%s:0", domain, filepath.Join("internal", "router", "injection.go")))
}

func generateMockRepository(domain string) error {
	domainFileName := strcase.ToSnake(domain)

	source := fmt.Sprintf("-source=%s.go", domainFileName)
	destination := fmt.Sprintf("-destination=%s", filepath.Join("mockrepository", domainFileName+".go"))

	cmd := exec.Command("mockgen", source, destination, "-package=mockrepository")
	cmd.Dir = filepath.Join("internal", "repository")
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	if string(stdout) != "" {
		fmt.Println(string(stdout))
	}

	return nil
}

func generateMockService(domain string) error {
	domainFileName := strcase.ToSnake(domain)

	source := fmt.Sprintf("-source=%s.go", domainFileName)
	destination := fmt.Sprintf("-destination=%s", filepath.Join("mockservice", domainFileName+".go"))

	cmd := exec.Command("mockgen", source, destination, "-package=mockservice")
	cmd.Dir = filepath.Join("internal", "service")
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	if string(stdout) != "" {
		fmt.Println(string(stdout))
	}

	return nil
}

func generateControllerTemplate(domain string, varErr1 string) error {
	controllerTemplate, err := template.GetControllerTemplate(domain, varErr1)
	if err != nil {
		return fmt.Errorf("err get controller template: %v", err)
	}

	path := filepath.Join("internal", "controller", "http", strcase.ToSnake(domain)+".go")
	err = os.WriteFile(path, []byte(controllerTemplate), 0666)
	if err != nil {
		return fmt.Errorf("err write controller template: %v", err)
	}

	return nil
}

func generateServiceTemplate(domain string, varErr1 string, varErr2 string) error {
	serviceTemplate, err := template.GetServiceTemplate(domain, varErr1, varErr2)
	if err != nil {
		return fmt.Errorf("err get service template: %v", err)
	}

	path := filepath.Join("internal", "service", strcase.ToSnake(domain)+".go")
	err = os.WriteFile(path, []byte(serviceTemplate), 0666)
	if err != nil {
		return fmt.Errorf("err write service template: %v", err)
	}

	return nil
}

func generateServiceTestTemplate(domain string, varErr1 string, varErr2 string) error {
	serviceTestTemplate, err := template.GteServiceTestTemplate(domain, varErr1, varErr2)
	if err != nil {
		return fmt.Errorf("err get service test template: %v", err)
	}

	path := filepath.Join("internal", "service", strcase.ToSnake(domain)+"_test.go")
	err = os.WriteFile(path, []byte(serviceTestTemplate), 0666)
	if err != nil {
		return fmt.Errorf("err write service test template: %v", err)
	}

	return nil
}

func generateRepositoryTemplate(domain string, varErr1 string, varErr2 string) error {
	repositoryTemplate, err := template.GetRepositoryTemplate(domain, varErr1, varErr2)
	if err != nil {
		return fmt.Errorf("err get repository template: %v", err)
	}

	path := filepath.Join("internal", "repository", strcase.ToSnake(domain)+".go")
	err = os.WriteFile(path, []byte(repositoryTemplate), 0666)
	if err != nil {
		return fmt.Errorf("err write repository template: %v", err)
	}

	return nil
}

func generateInjectionTemplate(domain string) error {
	path := filepath.Join("internal", "router", "injection.go")
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("err open injection.go file: %err", err)
		}

		if err = generateNewInjectionTemplate(domain, path); err != nil {
			return fmt.Errorf("err generate new injection template: %v", err)
		}

		return nil
	}
	defer f.Close()

	injectionAppendTemplate, err := template.GetInjectionAppendTemplate(domain)
	if err != nil {
		return fmt.Errorf("err append injection: %v", err)
	}

	if _, err := f.WriteString(injectionAppendTemplate); err != nil {
		return fmt.Errorf("err write injection append: %v", err)
	}

	return nil
}

func generateNewInjectionTemplate(domain string, path string) error {
	injectionTemplate, err := template.GetInjectionTemplate(domain)
	if err != nil {
		return fmt.Errorf("err get injection template: %v", err)
	}
	err = os.WriteFile(path, []byte(injectionTemplate), 0666)
	if err != nil {
		return fmt.Errorf("err write injection template: %v", err)
	}
	return nil
}
