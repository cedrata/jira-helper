package cmd

import (
	"errors"
	"os"

	"github.com/cedrata/jira-helper/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string
	conf *config.Config

	rootCmd = &cobra.Command{
		Use:   "jirahelper <subcommand>",
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")

	for i := range cmds {
		rootCmd.AddCommand(&cmds[i])
	}
}

func initConfig() {
	var err error 

	if cfgFile != "" {
		// Use config file from the flag.
		conf, err = config.LoadLocalConfig(cfgFile, ".jira.config")
		cobra.CheckErr(err)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		conf, err = config.LoadLocalConfig(home, ".jira.config")
		cobra.CheckErr(err)
	}
}

var cmds = []cobra.Command{
	{
		Use:   "issues <userId>",
		Short: "Get issues for a user",
		Long:  `This create a new agile documentation directory`,
		RunE:  getStory,
	},
}

func getStory(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("expected one element")
	}
	return nil
}
