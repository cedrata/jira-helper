package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/cedrata/jira-helper/pkg/config"
	"github.com/cedrata/jira-helper/pkg/jira"
	"github.com/cedrata/jira-helper/pkg/markdown"
	"github.com/cedrata/jira-helper/pkg/rest"
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

	newDocCmd = &cobra.Command{
		Use:   "doc [flags] <dir> <story id>",
		Short: "Create a new doc directory",
		Long: `This create a new agile documentation directory with a 
        presentation and a documentation markdown file`,
		RunE: writeStoryTemplate,
	}

	testsCmd = &cobra.Command{
		Use:   "tests [flags] <file>",
		Short: "Create a test table with all tests",
		Long:  `This create a test table in markdown with all tests`,
		RunE:  writeTestList,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("host", "H", "", "jira instance host")
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))

	rootCmd.PersistentFlags().StringP("token", "t", "", "jira instance token")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))

	rootCmd.PersistentFlags().StringP("project", "p", "", "jira project name")
	viper.BindPFlag("project", rootCmd.PersistentFlags().Lookup("project"))

	// TODO: to move in different files
	newDocCmd.Flags().StringP("user", "u", "AF82260", "user name to filter issues for")
	newDocCmd.Flags().BoolP("active-sprint", "a", true, "select the issues only in active sprints")
	rootCmd.AddCommand(newDocCmd)

	testsCmd.Flags().String("type", "test", "select the issue type")
	rootCmd.AddCommand(testsCmd)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/.jhelp.config)"))
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
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

func writeStoryTemplate(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("expected  2 args found %d", len(args))
	}

	dir := args[0]
	id := args[1]

	resp, err := rest.Get(rest.GetIssues, http.DefaultClient, viper.GetViper())
	if err != nil {
		return err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(*resp, &m)
	if err != nil {
		return err
	}

	issues := extractIssuesMap(m)
	story, ok := issues[id]
	if !ok {
		return fmt.Errorf("error key %s not found", args[0])
	}

	err = os.Mkdir(dir, 0o777)
	if err != nil {
		return err
	}

	err = markdown.WriteStub(story, path.Join(dir, markdown.DocFileSrc),
		markdown.DocTemplate)
	if err != nil {
		return err
	}

	err = markdown.WriteStub(story, path.Join(dir, markdown.PresFileSrc),
		markdown.PresTemplate)
	if err != nil {
		return err
	}

	fmt.Printf("created doc dir %s\n", args[0])
	return nil
}

func writeTestList(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("expected  1 args found %d", len(args))
	}

	file := args[0]

	resp, err := rest.Get(rest.GetIssues, http.DefaultClient, viper.GetViper())
	if err != nil {
		return err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(*resp, &m)
	if err != nil {
		return err
	}

	issues := extractIssues(m)

	testList := jira.TestList{Tests: issues}
	err = markdown.WriteTestTable(testList, file)
	if err != nil {
		return err
	}

	fmt.Printf("created test list %s\n", args[0])
	return nil
}
