package configure

import (
	"github.com/cedrata/jira-helper/cmd/configure/createprofile"
	"github.com/spf13/cobra"
)

var ConfigureCmd *cobra.Command

func init() {
	ConfigureCmd = &cobra.Command{
		Use: "configure",
	}

	ConfigureCmd.AddCommand(createprofile.CreateProfileCmd)
}
