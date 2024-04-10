package configure

import (
	"github.com/cedrata/jira-helper/cmd/configure/deleteprofile"
	"github.com/spf13/cobra"
)

var DeleteCmd *cobra.Command

func init() {
	DeleteCmd = &cobra.Command{
		Use: "delete",
	}

	DeleteCmd.AddCommand(deleteprofile.DeleteProfileCmd)
}
