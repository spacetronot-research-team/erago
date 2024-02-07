package createproject

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/erago/cmd/createdomain"
	"github.com/spacetronot-research-team/erago/cmd/createproject/template"
	"github.com/spacetronot-research-team/erago/common/gomod"
)

// CreateProject is main func to create new project.
func CreateProject(projectName string, moduleName string) {
	logrus.Info("create project start")

	logrus.Info("create project dir")
	projectPath := filepath.Join(".", projectName)
	if err := os.MkdirAll(projectPath, os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir projectPath: %v", err))
	}

	logrus.Info("run go mod init in project dir")
	if err := runGoModInit(moduleName, projectPath); err != nil {
		logrus.Fatal(fmt.Errorf("err run go mod init in project path: %v", err))
	}

	logrus.Info("create internal/controller/http dir in project dir")
	if err := os.MkdirAll(filepath.Join(projectPath, "internal", "controller", "http"), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir projectPath/internal/controller/http/: %v", err))
	}

	logrus.Info("create internal/service dir in project dir")
	if err := os.MkdirAll(filepath.Join(projectPath, "internal", "service"), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir projectPath/internal/service: %v", err))
	}

	logrus.Info("create internal/repository dir in project dir")
	if err := os.MkdirAll(filepath.Join(projectPath, "internal", "repository"), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir projectPath/internal/repository: %v", err))
	}

	logrus.Info("create internal/router dir in project dir")
	if err := os.MkdirAll(filepath.Join(projectPath, "internal", "router"), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir projectPath/internal/router: %v", err))
	}

	logrus.Info("create database dir in project dir")
	if err := os.MkdirAll(filepath.Join(projectPath, "database"), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir projectPath/database: %v", err))
	}

	logrus.Info("create open.go file in database dir")
	if err := generateDatabaseOpenTemplate(projectPath); err != nil {
		logrus.Fatal(fmt.Errorf("err write file projectPath/database/open.go: %v", err))
	}

	logrus.Info("create cmd dir in project dir")
	if err := os.MkdirAll(filepath.Join(projectPath, "cmd"), os.ModePerm); err != nil {
		logrus.Fatal(fmt.Errorf("err mkdir projectPath/cmd: %v", err))
	}

	logrus.Info("create main.go file in cmd dir")
	if err := generateMainTemplate(projectPath, moduleName); err != nil {
		logrus.Fatal(fmt.Errorf("err write file projectPath/cmd/main.go: %v", err))
	}

	logrus.Info("change current dir to project dir")
	if err := os.Chdir(projectPath); err != nil {
		logrus.Fatal(fmt.Errorf("err change work dir to projectPath: %v", err))
	}

	logrus.Info("create domain hello world")
	createdomain.CreateDomain("hello world")

	logrus.Info("run go mod tidy")
	if err := gomod.RunGoModTidy(); err != nil {
		logrus.Fatal(fmt.Errorf("err run go mod tidy: %v", err))
	}

	logrus.Info("create project finish, go to your project:\n\tcd ", projectPath)
}

func runGoModInit(moduleName string, projectPath string) error {
	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Dir = projectPath
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	if string(stdout) != "" {
		fmt.Println(string(stdout))
	}

	return nil
}

func generateDatabaseOpenTemplate(projectPath string) error {
	databaseOpenTemplate := template.GetDatabaseOpenTemplate()

	path := filepath.Join(projectPath, "database", "open.go")
	if err := os.WriteFile(path, []byte(databaseOpenTemplate), 0666); err != nil {
		return fmt.Errorf("err write database open template: %v", err)
	}

	return nil
}

func generateMainTemplate(projectPath string, moduleName string) error {
	mainTemplate, err := template.GetMainTemplate(moduleName)
	if err != nil {
		return fmt.Errorf("err get main template: %v", err)
	}

	path := filepath.Join(projectPath, "cmd", "main.go")
	if err := os.WriteFile(path, []byte(mainTemplate), os.ModePerm); err != nil {
		return fmt.Errorf("err write main.go template: %v", err)
	}

	return nil
}
