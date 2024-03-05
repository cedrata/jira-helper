package cmd

import (
	"fmt"
	"os"

	"github.com/cedrata/jira-helper/cmd/issues"
	"github.com/cedrata/jira-helper/cmd/issuesearch"
	"github.com/cedrata/jira-helper/cmd/myself"
	"github.com/cedrata/jira-helper/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var v *viper.Viper

var rootCmd = &cobra.Command{
	Use:   "jhelp [flags] <command> ",
	Short: "An helper for using JIRA on CLI",
	Long:  `An helper for using JIRA on CLI`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// When adding the 'config' subbcomamnd ignore validation of the config
		// and execute no matter what.

		// In case of any other command being invoked then the validation is
		// enabled

		// If 'config' is not at index 1 in the cmd.CommandPath() after
		// a string split then it's not a configuration command
		//
		// fmt.Printf("PersistentPreRun: %s\n", cmd.CommandPath())

		// IT WORKS :-)
		var err error
		config.ConfigData, err = config.ValidateProfile(v.GetString("profile"), v)
		if err != nil {
			cobra.CheckErr(err)
		}

        fmt.Println(config.ConfigData.Host, config.ConfigData.Token)
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("host", "H", "", "jira instance host")
	rootCmd.PersistentFlags().StringP("token", "t", "", "jira instance token")
	rootCmd.PersistentFlags().StringP("profile", "p", "default", "configuration profile")
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jhelp.config)")

	_ = viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	_ = viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	_ = viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("profile"))

	rootCmd.AddCommand(issues.IssuesCmd)
	rootCmd.AddCommand(issuesearch.IssueSearchCmd)
	rootCmd.AddCommand(myself.MyselfCmd)

	v = viper.GetViper()
}

func initConfig() {
	var err error
	var configPath string
	const configName = config.DefaultConfigName

	configPath, err = os.UserHomeDir()
	if err != nil {
		cobra.CheckErr(err)
	}

	err = config.LoadLocalConfig(configPath, configName, v)
	if err != nil {
		cobra.CheckErr(err)
	}

}
