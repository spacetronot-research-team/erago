package cli

import (
	"fmt"

	"github.com/spacetronot-research-team/erago/internal/service"
	"github.com/spf13/cobra"
)

type ExplainController struct {
	explainService service.Explain
}

func NewExplainController(explainService service.Explain) *ExplainController {
	return &ExplainController{
		explainService: explainService,
	}
}

func (ec *ExplainController) GetExplainCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:     "explain",
		Short:   "Explain code architecture",
		Long:    "Explain code architecture.",
		Args:    cobra.MaximumNArgs(0),
		Run:     ec.printCodeArchExplaination,
		Example: "erago explain",
	}
	return versionCmd
}

// printCodeArchExplaination print code arch explanation to standard output.
func (ec *ExplainController) printCodeArchExplaination(cmd *cobra.Command, args []string) {
	codeArchExplanation := ec.explainService.GetCodeArchExplanation()
	fmt.Println(codeArchExplanation) //nolint:forbidigo
}
