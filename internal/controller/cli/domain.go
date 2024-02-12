package cli

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/erago/internal/service"
	"github.com/spacetronot-research-team/erago/pkg/ctxutil"
	"github.com/spacetronot-research-team/erago/pkg/gomod"
	"github.com/spf13/cobra"
)

type DomainController struct {
	domainService service.Domain
}

func NewDomainController(domainService service.Domain) *DomainController {
	return &DomainController{
		domainService: domainService,
	}
}

func (dc *DomainController) GetCreateDomainCmd() *cobra.Command {
	createDomainCmd := &cobra.Command{
		Use:     "create-domain [domain]",
		Short:   "Create new domain with the provided domain name",
		Long:    "Create new domain with the provided domain name. If no domain name is provided, it defaults to 'hello world'.", //nolint:lll
		Args:    cobra.MaximumNArgs(1),
		Run:     dc.CreateDomain,
		Example: "erago create-domain profile",
	}
	return createDomainCmd
}

// CreateDomain create new domain.
func (dc *DomainController) CreateDomain(cmd *cobra.Command, args []string) {
	domain := "hello world"

	if len(args) > 0 {
		domain = args[0]
	}

	ctx := ctxutil.WithDomain(context.Background(), domain)

	moduleName, err := gomod.GetModuleName()
	if err != nil {
		logrus.Fatal(err)
	}
	ctx = ctxutil.WithModuleName(ctx, moduleName)

	projectPath, err := os.Getwd()
	if err != nil {
		logrus.Fatal(err)
	}
	ctx = ctxutil.SetAllDirPath(ctx, projectPath)

	projectName, err := gomod.GetProjectNameFromGoMod()
	if err != nil {
		logrus.Fatal(err)
	}
	ctx = ctxutil.WithProjectName(ctx, projectName)

	dc.domainService.CreateDomain(ctx)
}
