package cmd

import (
	"fmt"

	"github.com/cedrata/jira-helper/pkg/config"
	"github.com/cedrata/jira-helper/pkg/jira"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/.jhelp.config)"))

	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("project", rootCmd.PersistentFlags().Lookup("project"))
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}

func initConfig() {
	var err error
	var configFile string

	configFile = viper.GetString("config")
	err = config.LoadLocalConfig(configFile, viper.GetViper())

	// If the configuration file is not provided and the default configuration
	// does not exists then the flag values are used
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok && configFile == "" {
		cobra.CheckErr(err)
	}
}

func extractIssue(issue interface{}) jira.Issue {
	var newIssue jira.Issue

	fields := issue.(map[string]interface{})["fields"].(map[string]interface{})
	newIssue.Key = issue.(map[string]interface{})["key"].(string)

	if temp, ok := fields["assignee"].(map[string]interface{}); ok {
		newIssue.Assignee = temp["name"].(string)
	}

	if temp, ok := fields["description"].(string); ok {
		newIssue.Description = temp
	}

	if temp, ok := fields["status"].(map[string]interface{}); ok {
		newIssue.Status = temp["name"].(string)
	}

	if temp, ok := fields["summary"].(string); ok {
		newIssue.Summary = temp
	}

	return newIssue
}

func extractIssues(result map[string]interface{}) []jira.Issue {
	var issues []interface{}
	var res []jira.Issue
	issues = result["issues"].([]interface{})

	for _, k := range issues {
		res = append(res, extractIssue(k))
	}

	return res
}

func extractIssuesMap(result map[string]interface{}) map[string]jira.Issue {
	var issues []interface{}

	res := make(map[string]jira.Issue)
	issues = result["issues"].([]interface{})

	for k := range issues {
		issue := extractIssue(k)
		res[issue.Key] = issue
	}

	return res
}

