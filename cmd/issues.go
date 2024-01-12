package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cedrata/jira-helper/pkg/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var issuesCmd *cobra.Command

func init() {
	issuesCmd = &cobra.Command{
		Use:   "issues [flags]",
		Short: "Get issues for a user and status",
		Long:  `This returns the stories that matches the applied filters`,
		RunE:  getStory,
	}

	issuesCmd.Flags().StringP("user", "u", "", "user name to filter issues for")
	viper.BindPFlag("user", issuesCmd.Flags().Lookup("user"))

	issuesCmd.Flags().StringP("status", "s", "", "jira status to filter for")
	viper.BindPFlag("status", issuesCmd.Flags().Lookup("status"))

	issuesCmd.Flags().BoolP("active-sprint", "a", true, "select the issues only in active sprints")
	viper.BindPFlag("active-sprint", issuesCmd.Flags().Lookup("active-sprint"))

	rootCmd.AddCommand(issuesCmd)
}

func getStory(cmd *cobra.Command, args []string) error {
	resp, err := rest.Get(rest.GetIssues, http.DefaultClient, viper.GetViper())
	if err != nil {
		return err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(*resp, &m)
	if err != nil {
		return err
	}

	fmt.Println(extractIssues(m))
	return nil
}
