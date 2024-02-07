package gomod

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunGoModTidy() error {
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

func GetModuleName() (string, error) {
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

func GetProjectNameFromModule() (string, error) {
	moduleName, err := GetModuleName()
	if err != nil {
		return "", fmt.Errorf("err get module name: %v", err)
	}

	moduleNameSplitted := strings.Split(moduleName, "/")
	lastIndex := len(moduleNameSplitted) - 1
	return moduleNameSplitted[lastIndex], nil
}
