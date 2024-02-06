package createproject

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spacetronot-research-team/erago/cmd/createdomain"
)

// CreateProject is main func to create new project.
func CreateProject(projectName string, moduleName string) {
	log.Println("create project start")

	log.Println("create project dir")
	projectPath := filepath.Join(".", projectName)
	if err := os.MkdirAll(projectPath, os.ModePerm); err != nil {
		log.Fatal(fmt.Errorf("err mkdir projectPath: %v", err))
	}

	log.Println("run go mod init in project dir")
	if err := runGoModInit(moduleName, projectPath); err != nil {
		log.Fatal(fmt.Errorf("err run go mod init in project path: %v", err))
	}

	log.Println("create internal/controller dir in project dir")
	if err := os.MkdirAll(filepath.Join(projectPath, "internal", "controller"), os.ModePerm); err != nil {
		log.Fatal(fmt.Errorf("err mkdir projectPath/internal/controller: %v", err))
	}

	log.Println("create internal/service dir in project dir")
	if err := os.MkdirAll(filepath.Join(projectPath, "internal", "service"), os.ModePerm); err != nil {
		log.Fatal(fmt.Errorf("err mkdir projectPath/internal/service: %v", err))
	}

	log.Println("create internal/repository dir in project dir")
	if err := os.MkdirAll(filepath.Join(projectPath, "internal", "repository"), os.ModePerm); err != nil {
		log.Fatal(fmt.Errorf("err mkdir projectPath/internal/repository: %v", err))
	}

	log.Println("create cmd dir in project dir")
	if err := os.MkdirAll(filepath.Join(projectPath, "cmd"), os.ModePerm); err != nil {
		log.Fatal(fmt.Errorf("err mkdir projectPath/cmd: %v", err))
	}

	log.Println("create main.go file in cmd dir")
	if err := os.WriteFile(filepath.Join(projectPath, "cmd", "main.go"), []byte("package main\n"), os.ModePerm); err != nil {
		log.Fatal(fmt.Errorf("err write file projectPath/cmd/main.go: %v", err))
	}

	log.Println("change current dir to project dir")
	if err := os.Chdir(projectPath); err != nil {
		log.Fatal(fmt.Errorf("err change work dir to projectPath: %v", err))
	}

	log.Println("create domain hello world")
	createdomain.CreateDomain("hello world")

	log.Println("run go mod tidy")
	if err := runGoModTidy(); err != nil {
		log.Fatal(fmt.Errorf("err run go mod tidy in projectPath: %v", err))
	}

	log.Println("create project finish, go to your project:\n\tcd", projectPath)
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

func runGoModTidy() error {
	cmd := exec.Command("go", "mod", "tidy")
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	if string(stdout) != "" {
		fmt.Println(string(stdout))
	}

	return nil
}
