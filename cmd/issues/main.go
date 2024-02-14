package issues

import (
	"github.com/spf13/cobra"
)

var IssuesCmd *cobra.Command

func init() {
	IssuesCmd = &cobra.Command{
		Use: "issues",
	}

	IssuesCmd.AddCommand(getTransitionsCmd)
}
