package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/erago/internal/router"
	"github.com/spf13/cobra"
)

func main() {
	if os.Getenv("ERAGO_DEBUG") == "true" {
		logrus.SetReportCaller(true)
	}

	rootCmd := &cobra.Command{
		Use:               "erago",
		Short:             "Erajaya CLI generate project",
		Long:              "Erajaya CLI generate project.",
		CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
	}

	router.RegisterCommand(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
