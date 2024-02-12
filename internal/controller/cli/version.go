package cli

import (
	"fmt"

	"github.com/spacetronot-research-team/erago/internal/service"
	"github.com/spf13/cobra"
)

type VersionController struct {
	versionService service.Version
}

func NewVersionController(versionService service.Version) *VersionController {
	return &VersionController{
		versionService: versionService,
	}
}

func (vc *VersionController) GetVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:     "version",
		Short:   "Print erago version",
		Long:    "Print erago version.",
		Args:    cobra.MaximumNArgs(0),
		Run:     vc.printCurrentVersion,
		Example: "erago version",
	}
	return versionCmd
}

// printCurrentVersion print current version to standard output.
func (vc *VersionController) printCurrentVersion(cmd *cobra.Command, args []string) { //nolint:all
	currentVersion := vc.versionService.GetCurrentVersion()
	fmt.Println(currentVersion) //nolint:all
}
