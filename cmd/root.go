package cmd

import (
	"github.com/cedrata/jira-helper/cmd/issues"
	"github.com/cedrata/jira-helper/cmd/issuesearch"
	"github.com/cedrata/jira-helper/cmd/myself"
	"github.com/cedrata/jira-helper/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "jhelp [flags] <command> ",
		Short: "An helper for using JIRA on CLI",
		Long:  `An helper for using JIRA on CLI`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("host", "H", "", "jira instance host")
	rootCmd.PersistentFlags().StringP("token", "t", "", "jira instance token")
	rootCmd.PersistentFlags().StringP("project", "p", "", "jira project name")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jhelp.config)")

	_ = viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	_ = viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	_ = viper.BindPFlag("project", rootCmd.PersistentFlags().Lookup("project"))
	_ = viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

    rootCmd.AddCommand(issues.IssuesCmd)
    rootCmd.AddCommand(issuesearch.IssueSearchCmd)
    rootCmd.AddCommand(myself.MyselfCmd)
}

func initConfig() {
	var err error
	configFile := viper.GetString("config")
	err = config.LoadLocalConfig(configFile, viper.GetViper())

	// If the configuration file is not provided and the default configuration
	// does not exists then the flag values are used
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok && configFile == "" {
		cobra.CheckErr(err)
	}
}
