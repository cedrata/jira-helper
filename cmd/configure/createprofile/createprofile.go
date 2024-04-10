package createprofile

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CreateProfileCmd *cobra.Command

func init() {
	CreateProfileCmd = &cobra.Command{
		Use:   "create-profile",
		Short: "Create profile",
		Long:  "Creates a new profile. Use this to create a profile with a name, host, and token.",
		RunE:  createProfileHandler,
	}

	_ = CreateProfileCmd.Flags().String("name", "", "Name of profile.")
	_ = CreateProfileCmd.Flags().String("host", "", "Host for profile.")
	_ = CreateProfileCmd.Flags().String("token", "", "Token for profile.")

	// Required flags
	_ = CreateProfileCmd.MarkFlagRequired("name")
	_ = CreateProfileCmd.MarkFlagRequired("host")
	_ = CreateProfileCmd.MarkFlagRequired("token")
}

func createProfileHandler(cmd *cobra.Command, args []string) error {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	host, err := cmd.Flags().GetString("host")
	if err != nil {
		return err
	}

	token, err := cmd.Flags().GetString("token")
	if err != nil {
		return err
	}

	v := viper.GetViper()

	if v == nil {
		return fmt.Errorf("unable to access viper instance")
	}

	newProfile := map[string]interface{}{
		"host":  host,
		"token": token,
	}

	if v.Sub("profiles").InConfig(name) {
		return fmt.Errorf("profile %s already exists", name)
	}

	viper.Set(fmt.Sprintf("profiles.%s", name), newProfile)

	return viper.WriteConfig()
}
