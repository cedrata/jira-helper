package selectprofile

import (
	"github.com/spf13/cobra"
)

var SelectProfileCmd *cobra.Command

func init() {
	SelectProfileCmd = &cobra.Command{
		Use:   "select-profile",
		Short: "Select profile",
		Long:  "Selects a profile based on provided profile name.",
		RunE:  selectProfileHandler,
	}

	_ = SelectProfileCmd.Flags().String("name", "", "Name of profile")
	_ = SelectProfileCmd.Flags().String("default", "", "Set to true when setting the default profile")
	_ = SelectProfileCmd.MarkFlagRequired("name")
}

func selectProfileHandler(cmd *cobra.Command, args []string) error {
	return nil
}
