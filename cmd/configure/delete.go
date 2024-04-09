package configure

import (
	"github.com/spf13/cobra"
)

var DeleteCmd *cobra.Command

func init() {
	DeleteCmd = &cobra.Command{
		Use: "configure",
	}

	// DeleteCmd.AddCommand(deleteprofile.DeleteProfileCmd)
}
