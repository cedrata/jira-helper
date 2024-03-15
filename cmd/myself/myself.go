package myself

import (
	"github.com/cedrata/jira-helper/cmd/myself/getcurrentuser"
	"github.com/spf13/cobra"
)

var MyselfCmd *cobra.Command

func init() {
	MyselfCmd = &cobra.Command{
		Use: "myself",
	}

	MyselfCmd.AddCommand(getcurrentuser.GetCurrentUserCmd)
}
