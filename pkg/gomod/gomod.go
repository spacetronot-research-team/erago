package gomod

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunGoModTidy() error {
	cmd := exec.Command("go", "mod", "tidy")
	stdout, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("err run go mod tidy: %v", err)
	}

	if string(stdout) != "" {
		fmt.Println(string(stdout)) //nolint:all
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
		expectedLenLine := 2
		if len(parts) >= expectedLenLine {
			moduleName := parts[1]
			return moduleName, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("err scan go.mod: %v", err)
	}

	return "", fmt.Errorf("err module name not found in go.mod file: %v", err)
}

func GetProjectNameFromGoMod() (string, error) {
	moduleName, err := GetModuleName()
	if err != nil {
		return "", fmt.Errorf("err get go.mod: %v", err)
	}

	moduleNameSplitted := strings.Split(moduleName, "/")
	lastIndex := len(moduleNameSplitted) - 1
	return moduleNameSplitted[lastIndex], nil
}
