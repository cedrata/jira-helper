package cmd

import (
	"os"

	"github.com/cedrata/jira-helper/cmd/configure"
	"github.com/cedrata/jira-helper/cmd/issues"
	"github.com/cedrata/jira-helper/cmd/issuesearch"
	"github.com/cedrata/jira-helper/cmd/myself"
	"github.com/cedrata/jira-helper/pkg/config"
	"github.com/cedrata/jira-helper/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var v *viper.Viper

var rootCmd *cobra.Command

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd = &cobra.Command{
		Use:               "jira-helper [flags] <command> ",
		Short:             "An helper for using JIRA on CLI",
		Long:              `An helper for using JIRA on CLI`,
		PersistentPreRunE: persistentPreRunHandler,
	}
	rootCmd.PersistentFlags().StringP("host", "H", "", "jira instance host")
	rootCmd.PersistentFlags().StringP("token", "t", "", "jira instance token")
	rootCmd.PersistentFlags().StringP("profile", "p", "default", "configuration profile")

	rootCmd.AddCommand(issues.IssuesCmd)
	rootCmd.AddCommand(issuesearch.IssueSearchCmd)
	rootCmd.AddCommand(myself.MyselfCmd)
	rootCmd.AddCommand(configure.ConfigureCmd)

	v = viper.GetViper()
}

func persistentPreRunHandler(cmd *cobra.Command, args []string) error {
	if cmd.Name() == "set" {
		return nil
	}

	var err error
	var configPath string
	const configName = config.DefaultConfigName

	config.ConfigData = &config.Config{}

	if configPath, err = os.UserHomeDir(); err != nil {
		return err
	}

	err = config.LoadLocalConfig(configPath, configName, v)
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok && err != nil {
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
