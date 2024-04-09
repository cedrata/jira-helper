package createprofile

import (
	"os"

	"github.com/spf13/cobra"
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
}

func createProfileHandler(cmd *cobra.Command, args []string) error {
	var err error
	var name string
	var host string
	var token string

	f, err := os.OpenFile("$HOME/.jira-helper.config", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("[" + name + "]\n")
	if err != nil {
		return err
	}

	_, err = f.WriteString("host=" + host + "\n")
	if err != nil {
		return err
	}

	_, err = f.WriteString("token=" + token + "\n")
	if err != nil {
		return err
	}

	return nil
}
