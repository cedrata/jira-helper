package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/cedrata/jira-helper/pkg/config"
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

	cmds = []cobra.Command{}
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
	cmds = append(cmds, *issuesCmd)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jhelp.config)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")

	for i := range cmds {
		rootCmd.AddCommand(&cmds[i])
	}
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

func extractIssues(result map[string]interface{}) []issue {
	var issues []interface{}
	var res []issue
	issues = result["issues"].([]interface{})

	for k := range issues {
		fields := issues[k].(map[string]interface{})["fields"].(map[string]interface{})
		key := issues[k].(map[string]interface{})["key"].(string)
		assignee := fields["assignee"].(map[string]interface{})["name"].(string)
		description := fields["description"].(string)
		status := fields["status"].(map[string]interface{})["name"].(string)
		res = append(res, issue{key, assignee, description, status})
	}

	return res
}

type issue struct {
	key         string
	assignee    string
	descritpion string
	status      string
}

func (i issue) String() string {
	return fmt.Sprintf(
		"\nkey: %s\nassignee: %s\nstatus: %s\ndescription: %s\n",
		i.key,
		i.assignee,
		i.status,
		i.descritpion,
	)
}
