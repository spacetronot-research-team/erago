package createdomain

import (
	"github.com/spf13/cobra"
)

// CLI is cobra command used for create new domain.
func CLI(cmd *cobra.Command, args []string) {
	domain := "hello world"

	if len(args) > 0 {
		domain = args[0]
	}

	CreateDomain(domain)
}
