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
	conf        *config.Config

	rootCmd = &cobra.Command{
		Use:   "jhelp <command>",
		Short: "An helper for using JIRA on CLI",
		Long:  `An helper for using JIRA on CLI`,
	}

	issuesCmd = &cobra.Command{
		Use:   "issues [options[--user=<user> --status=<status>]] --project=<project>",
		Short: "Get issues for a user and status",
		Long:  `This create a new agile documentation directory`,
		RunE:  getStory,
	}

	newDocCmd = &cobra.Command{
		Use:   "doc <dir> <story id>",
		Short: "Create a new doc directory",
		Long:  `This create a new agile documentation directory with a presentation and a documentation markdown file`,
		RunE:  writeStoryTemplate,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	initConfig()

	rootCmd.PersistentFlags().StringP("host", "H", conf.JiraUrl, "jira instance host")
	rootCmd.PersistentFlags().StringP("token", "t", conf.Token, "jira instance host")
	rootCmd.PersistentFlags().StringP("project", "p", conf.Project, "jira instance host")

	issuesCmd.Flags().StringP("user", "u", "", "user name to filter issues for")
	issuesCmd.Flags().StringP("status", "s", "", "jira status to filter fo")
	issuesCmd.Flags().BoolP("active-sprint", "a", true, "select the issues only in active sprints")
	rootCmd.AddCommand(issuesCmd)

	newDocCmd.Flags().StringP("user", "u", "AF82260", "user name to filter issues for")
	newDocCmd.Flags().BoolP("active-sprint", "a", true, "select the issues only in active sprints")
	rootCmd.AddCommand(newDocCmd)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jhelp.config)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")

}

func initConfig() {
	var err error

	if cfgFile != "" {
		// Use config file from the flag.
		conf, err = config.LoadLocalConfig(cfgFile, ".jhelp.config")
		cobra.CheckErr(err)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		conf, err = config.LoadLocalConfig(home, ".jhelp.config")
		cobra.CheckErr(err)
	}
}

// Get all stories for a project,
// those can be filtered by:
//
//	status
//	user
//
// a flag to select only issues from active sprint is available
func getStory(cmd *cobra.Command, args []string) error {
	resp, err := rest.Get(rest.GetIssues, http.DefaultClient, cmd.Flags())
	if err != nil {
		return err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(*resp, &m)
	if err != nil {
		return err
	}

	// fmt.Println(m["issues"])
	fmt.Println(extractIssues(m))
	return nil
}

func extractIssues(result map[string]interface{}) []jira.Issue {
	var issues []interface{}
	var res []jira.Issue
	issues = result["issues"].([]interface{})

	for k := range issues {
		fields := issues[k].(map[string]interface{})["fields"].(map[string]interface{})
		key := issues[k].(map[string]interface{})["key"].(string)
		assignee := fields["assignee"].(map[string]interface{})["name"].(string)
		description := fields["description"].(string)
		status := fields["status"].(map[string]interface{})["name"].(string)
		summary := fields["summary"].(string)
		res = append(res, jira.Issue{
			Key:         key,
			Assignee:    assignee,
			Description: description,
			Status:      status,
			Summary:     summary,
		})
	}

	return res
}

func extractIssuesMap(result map[string]interface{}) map[string]jira.Issue {
	var issues []interface{}

	res := make(map[string]jira.Issue)
	issues = result["issues"].([]interface{})

	for k := range issues {
		fields := issues[k].(map[string]interface{})["fields"].(map[string]interface{})
		key := issues[k].(map[string]interface{})["key"].(string)
		assignee := fields["assignee"].(map[string]interface{})["name"].(string)
		description := fields["description"].(string)
		status := fields["status"].(map[string]interface{})["name"].(string)
		summary := fields["summary"].(string)
		res[key] = jira.Issue{
			Key:         key,
			Assignee:    assignee,
			Description: description,
			Status:      status,
			Summary:     summary,
		}
	}

	return res
}

func writeStoryTemplate(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("expected  2 args found %d", len(args))
	}
	
	dir := args[0]
	id := args[1] 

	resp, err := rest.Get(rest.GetIssues, http.DefaultClient, cmd.Flags())
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
