package configure

import (
	"github.com/cedrata/jira-helper/cmd/configure/set"
	"github.com/spf13/cobra"
)

var ConfigureCmd *cobra.Command

func init() {
	ConfigureCmd = &cobra.Command{
		Use: "configure",
	}

	ConfigureCmd.AddCommand(set.SetCmd)
}
