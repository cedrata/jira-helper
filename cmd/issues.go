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
	issuesCmd.Flags().StringP("status", "s", "", "jira status to filter for")
	issuesCmd.Flags().StringP("output", "o", "", "store the result into the specified file")
	issuesCmd.Flags().BoolP("active-sprint", "a", true, "select the issues only in active sprints")

	viper.BindPFlag("user", issuesCmd.Flags().Lookup("user"))
	viper.BindPFlag("status", issuesCmd.Flags().Lookup("status"))
	viper.BindPFlag("output", issuesCmd.Flags().Lookup("output"))
	viper.BindPFlag("active-sprint", issuesCmd.Flags().Lookup("active-sprint"))

	rootCmd.AddCommand(issuesCmd)
}

// Add workflow to incorporate the `writeStoryTemplate` function into this one
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
