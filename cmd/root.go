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
		Use:   "issues [options[--user <user> --status <status>]] --project <project>",
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

	// Set default flags values from config file:
	// TO COMPLETE...
	// fmt.Println(conf.JiraUrl)
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
    // TODO:
    //  makr flags required
    //  store flags safely (array or struct somewhere, must be const)
    //  generate jql query to properly fetch and search for issues


    status, _ := cmd.Flags().GetString("status")
    fmt.Println(status)

	resp, err := rest.Get(rest.GetIssues, conf, http.DefaultClient, cmd.Flags())

	if err != nil {
		return err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(*resp, &m)
	if err != nil {
		return err
	}

	// fmt.Println(m["issues"])
	return nil
}
