package set

import (
	"fmt"
	"os"

	"github.com/cedrata/jira-helper/pkg/config"
	"github.com/cedrata/jira-helper/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	v      *viper.Viper
	SetCmd *cobra.Command
)

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

	config.ConfigData = &config.Config{}

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

	v.Set(fmt.Sprintf("%v.host", name), host)
	v.Set(fmt.Sprintf("%v.token", name), token)
	v.Set("profile", name)
	v.Set("host", host)
	v.Set("token", token)

	if err = v.WriteConfig(); err != nil {
		return err
	}

	profile := v.GetString("profile")
	err = v.UnmarshalKey(profile, config.ConfigData)
	if err != nil {
		return err
	}

	if token := v.GetString("token"); token != "" {
		config.ConfigData.Token = token
	}

	if host := v.GetString("host"); host != "" {
		config.ConfigData.Host = host
	}

	if err = utils.ValidateStruct(*config.ConfigData); err != nil {
		return err
	}

	return nil
}
