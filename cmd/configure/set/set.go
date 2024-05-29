package set

import (
	"fmt"
	"os"

	"github.com/cedrata/jira-helper/pkg/config"
	"github.com/cedrata/jira-helper/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var SetCmd *cobra.Command

func init() {
	SetCmd = &cobra.Command{
		Use:   "set",
		Short: "Set profile",
		Long:  "Set profile in configuration file. If provided profile does not exist, it will be created. If profile is set, the default profile is used. If provided profile already exists, it will be updated.",
		RunE:  setProfileHandler,
	}

	_ = SetCmd.Flags().String("name", "", "Name of profile.")
	_ = SetCmd.Flags().String("host", "", "Host for profile.")
	_ = SetCmd.Flags().String("token", "", "Token for profile.")

	// Required flags
	_ = SetCmd.MarkFlagRequired("name")
	_ = SetCmd.MarkFlagRequired("host")
	_ = SetCmd.MarkFlagRequired("token")
}

func setProfileHandler(cmd *cobra.Command, args []string) error {
	var err error
	var configPath string

	name, _ := cmd.Flags().GetString("name")
	host, _ := cmd.Flags().GetString("host")
	token, _ := cmd.Flags().GetString("token")

	v := viper.GetViper()

	if configPath, err = os.UserHomeDir(); err != nil {
		return err
	}

	err = config.LoadLocalConfig(configPath, config.DefaultConfigName, v)
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok && err != nil {
		return err
	}

	if name == "default" {
		v.SetDefault("host", host)
		v.SetDefault("token", token)
	} else {
		v.Set(fmt.Sprintf("%s.host", name), host)
		v.Set(fmt.Sprintf("%s.token", name), token)
	}

	if err = v.WriteConfigAs(configPath); err != nil {
		return err
	}

	if err = utils.ValidateStruct(config.Config{Host: host, Token: token}); err != nil {
		return err
	}

	return nil
}
