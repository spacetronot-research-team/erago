package createdomain

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/spacetronot-research-team/erago/cmd/createdomain/template"
)

// CreateDomain is main func to create new domain.
func CreateDomain(domain string) {
	moduleName, err := getModuleName()
	if err != nil {
		log.Fatal(err)
	}

	if err := generateControllerTemplate(domain, moduleName); err != nil {
		log.Fatal(err)
	}

	if err := generateServiceTemplate(domain, moduleName); err != nil {
		log.Fatal(err)
	}

	if err := generateRepositoryTemplate(domain); err != nil {
		log.Fatal(err)
	}
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

func generateControllerTemplate(domain string, moduleName string) error {
	controllerTemplate, err := template.GetControllerTemplate(domain, moduleName)
	if err != nil {
		return fmt.Errorf("err get controller template: %v", err)
	}

	path := fmt.Sprintf("internal/controller/%s.go", strcase.ToSnake(domain))
	err = os.WriteFile(path, []byte(controllerTemplate), 0666)
	if err != nil {
		return fmt.Errorf("err write controller template: %v", err)
	}

	return nil
}

func generateServiceTemplate(domain string, moduleName string) error {
	serviceTemplate, err := template.GetServiceTemplate(domain, moduleName)
	if err != nil {
		return fmt.Errorf("err get service template: %v", err)
	}

	path := fmt.Sprintf("internal/service/%s.go", strcase.ToSnake(domain))
	err = os.WriteFile(path, []byte(serviceTemplate), 0666)
	if err != nil {
		return fmt.Errorf("err write service template: %v", err)
	}

	return nil
}

func generateRepositoryTemplate(domain string) error {
	repositoryTemplate, err := template.GetRepositoryTemplate(domain)
	if err != nil {
		return fmt.Errorf("err get repository template: %v", err)
	}

	path := fmt.Sprintf("internal/repository/%s.go", strcase.ToSnake(domain))
	err = os.WriteFile(path, []byte(repositoryTemplate), 0666)
	if err != nil {
		return fmt.Errorf("err write repository template: %v", err)
	}

	return nil
}
