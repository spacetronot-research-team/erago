package main

import (
	"log"

	"github.com/spacetronot-research-team/erago/cmd/createdomain"
	"github.com/spacetronot-research-team/erago/cmd/createproject"
	"github.com/spacetronot-research-team/erago/cmd/explain"
	"github.com/spacetronot-research-team/erago/cmd/version"
	"github.com/spf13/cobra"
)

// Root cobra CLI.
func main() {
	rootCmd := &cobra.Command{
		Use:               "erago",
		Short:             "erago is Erajaya CLI generate project",
		Long:              "erago is Erajaya CLI generate project.",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	rootCmd.AddCommand(getCreateDomainCmd())
	rootCmd.AddCommand(getCreateProjectCmd())
	rootCmd.AddCommand(getVersionCmd())
	rootCmd.AddCommand(getExplainCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func getCreateDomainCmd() *cobra.Command {
	createDomainCmd := &cobra.Command{
		Use:   "create-domain [domain]",
		Short: "create-domain will create new domain with the provided domain name",
		Long:  "create-domain will create new domain with the provided domain name. If no domain name is provided, it defaults to 'hello world'.",
		Args:  cobra.MaximumNArgs(1),
		Run:   createdomain.CLI,
	}
	return createDomainCmd
}

func getCreateProjectCmd() *cobra.Command {
	createProjectCmd := &cobra.Command{
		Use:   "create-project [project-name] [module-name]",
		Short: "create-project will create new project with the provided domain name",
		Long:  "create-project will create new project with the provided domain name.",
		Args:  cobra.MinimumNArgs(2),
		Run:   createproject.CLI,
	}
	return createProjectCmd
}

func getVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "version will print erago version",
		Long:  "version will print erago version.",
		Args:  cobra.MaximumNArgs(0),
		Run:   version.CLI,
	}
	return versionCmd
}

func getExplainCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "explain",
		Short: "explain will explain code architecture",
		Long:  "explain will explain code architecture.",
		Args:  cobra.MaximumNArgs(0),
		Run:   explain.CLI,
	}
	return versionCmd
}
