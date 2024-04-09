package editprofile

import (
	"github.com/spf13/cobra"
)

var EditProfileCmd *cobra.Command

func init() {
	EditProfileCmd = &cobra.Command{
		Use:   "edit-profile",
		Short: "Edit profile",
		Long:  "Edit profile.",
		RunE:  editProfileHandler,
	}

	_ = EditProfileCmd.Flags().String("name", "", "Name of profile.")
	_ = EditProfileCmd.Flags().String("host", "", "Host for profile.")
	_ = EditProfileCmd.Flags().String("token", "", "Token for profile.")
}

func editProfileHandler(cmd *cobra.Command, args []string) error {
	return nil
}
