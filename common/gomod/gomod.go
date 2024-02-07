package gomod

import (
	"fmt"
	"os/exec"
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
