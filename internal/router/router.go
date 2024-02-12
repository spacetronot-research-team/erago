package router

import (
	"github.com/spf13/cobra"
)

func RegisterCommand(rootCmd *cobra.Command) {
	versionController := getVersionController()
	explainController := getExplainController()
	domainController := GetDomainController()
	projectController := getProjectController()

	rootCmd.AddCommand(versionController.GetVersionCmd())
	rootCmd.AddCommand(explainController.GetExplainCmd())
	rootCmd.AddCommand(domainController.GetCreateDomainCmd())
	rootCmd.AddCommand(projectController.GetCreateProjectCmd())
}
