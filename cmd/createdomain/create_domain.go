package createdomain

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/erago/cmd/createdomain/template"
	"github.com/spacetronot-research-team/erago/common/gomod"
)

// CreateDomain is main func to create new domain.
func CreateDomain(domain string) {
	logrus.Println("create domain start")

	logrus.Println("get module name")
	moduleName, err := getModuleName()
	if err != nil {
		logrus.Fatal(err)
	}

	varErr1 := fmt.Sprintf("%d", rand.Int())
	varErr2 := fmt.Sprintf("%d", rand.Int())

	logrus.Println("generate controller template")
	if err := generateControllerTemplate(domain, moduleName, varErr1); err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("generate service template")
	if err := generateServiceTemplate(domain, moduleName, varErr1, varErr2); err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("generate repository template")
	if err := generateRepositoryTemplate(domain, varErr1, varErr2); err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("generate injection")
	if err := generateInjectionTemplate(domain, moduleName); err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("generate mock repository using mockgen")
	if err := generateMockRepository(domain); err != nil {
		logrus.Println("err generate mock repository using mockgen, did you able to run `mockgen`?")
		logrus.Println("try to install mockgen:\n\tgo install go.uber.org/mock/mockgen@latest")
		logrus.Println("create domain finish without mock repository and service test template")
		return
	}

	logrus.Println("generate service test template")
	if err := generateServiceTestTemplate(domain, moduleName, varErr1, varErr2); err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("generate mock service using mockgen")
	if err := generateMockService(domain); err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("run go mod tidy")
	if err := gomod.RunGoModTidy(); err != nil {
		logrus.Fatal(fmt.Errorf("err run go mod tidy: %v", err))
	}

	logrus.Println("create domain finish")
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

func getModuleName() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", fmt.Errorf("err open go.mod: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		isLineStartsWithModule := strings.HasPrefix(line, "module ")
		if !isLineStartsWithModule {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) >= 2 {
			moduleName := parts[1]
			return moduleName, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scanner err: %v", err)
	}

	return "", errors.New("module name not found in go.mod file")
}

func generateControllerTemplate(domain string, moduleName string, varErr1 string) error {
	controllerTemplate, err := template.GetControllerTemplate(domain, moduleName, varErr1)
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

func generateServiceTemplate(domain string, moduleName string, varErr1 string, varErr2 string) error {
	serviceTemplate, err := template.GetServiceTemplate(domain, moduleName, varErr1, varErr2)
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

func generateServiceTestTemplate(domain string, moduleName string, varErr1 string, varErr2 string) error {
	serviceTestTemplate, err := template.GteServiceTestTemplate(domain, moduleName, varErr1, varErr2)
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

func generateInjectionTemplate(domain string, moduleName string) error {
	// TODO: check if file exists
	// TODO: if exists, append
	// TODO: if not, create new
	path := filepath.Join("internal", "router", "injection.go")
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("err open injection.go file: %err", err)
		}

		if err = generateNewInjectionTemplate(domain, moduleName, path); err != nil {
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

func generateNewInjectionTemplate(domain string, moduleName string, path string) error {
	injectionTemplate, err := template.GetInjectionTemplate(domain, moduleName)
	if err != nil {
		return fmt.Errorf("err get injection template: %v", err)
	}
	err = os.WriteFile(path, []byte(injectionTemplate), 0666)
	if err != nil {
		return fmt.Errorf("err write injection template: %v", err)
	}
	return nil
}
