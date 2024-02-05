package main

import (
	"log"

	"github.com/spacetronot-research-team/erago/cmd/createdomain"
	"github.com/spacetronot-research-team/erago/cmd/createproject"
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
