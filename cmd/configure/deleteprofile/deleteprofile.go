package deleteprofile

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	// Required flags
	_ = DeleteProfileCmd.MarkFlagRequired("name")
}

func deleteProfileHandler(cmd *cobra.Command, args []string) error {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	v := viper.GetViper()

	if v == nil {
		return fmt.Errorf("unable to access viper instance")
	}

	if !v.InConfig(name) {
		return fmt.Errorf("profile %s not found", name)
	}

	delete(viper.Get(name).(map[string]interface{}), name)

	fmt.Printf("Profile %s deleted successfully.\n", name)

	return viper.WriteConfig()
}
