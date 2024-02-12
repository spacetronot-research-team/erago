package cli

import (
	"context"
	"strings"

	"github.com/spacetronot-research-team/erago/internal/service"
	"github.com/spacetronot-research-team/erago/pkg/ctxutil"
	"github.com/spf13/cobra"
)

type ProjectController struct {
	projectService service.Project
}

func NewProjectController(projectService service.Project) *ProjectController {
	return &ProjectController{
		projectService: projectService,
	}
}

func (pc *ProjectController) GetCreateProjectCmd() *cobra.Command {
	createProjectCmd := &cobra.Command{
		Use:     "create-project [module-name]",
		Short:   "Create new project with the provided module name",
		Long:    "Create new project with the provided module name.",
		Args:    cobra.MinimumNArgs(1),
		Run:     pc.CreateProject,
		Example: "erago create-project github.com/eratech/go-customer",
	}
	return createProjectCmd
}

// CreateProject create new project.
func (pc *ProjectController) CreateProject(cmd *cobra.Command, args []string) {
	moduleName := args[0]

	ctx := ctxutil.WithDomain(context.Background(), "hello world")
	ctx = ctxutil.WithModuleName(ctx, moduleName)

	moduleNameSplitted := strings.Split(moduleName, "/")
	lastIndex := len(moduleNameSplitted) - 1
	projectName := moduleNameSplitted[lastIndex]
	ctx = ctxutil.WithProjectName(ctx, projectName)

	pc.projectService.CreateProject(ctx)
}
