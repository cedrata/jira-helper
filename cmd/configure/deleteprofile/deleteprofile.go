package deleteprofile

import (
	"github.com/spf13/cobra"
)

var DeleteProfileCmd *cobra.Command

func init() {
	DeleteProfileCmd = &cobra.Command{
		Use:   "delete-profile",
		Short: "Delete profile",
		Long:  "Deletes a profile based on provided profile name.",
		RunE:  deleteProfileHandler,
	}

	_ = DeleteProfileCmd.Flags().String("name", "", "Name of profile.")
}

func deleteProfileHandler(cmd *cobra.Command, args []string) error {
	return nil
}
