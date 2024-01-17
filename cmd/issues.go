package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cedrata/jira-helper/pkg/markdown"
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
	issuesCmd.Flags().StringP("output", "o", "", "store the result into a markdown file having the given output name")
	issuesCmd.Flags().BoolP("active-sprint", "a", true, "select the issues only in active sprints")

	viper.BindPFlag("user", issuesCmd.Flags().Lookup("user"))
	viper.BindPFlag("status", issuesCmd.Flags().Lookup("status"))
	viper.BindPFlag("output", issuesCmd.Flags().Lookup("output"))
	viper.BindPFlag("active-sprint", issuesCmd.Flags().Lookup("active-sprint"))

	rootCmd.AddCommand(issuesCmd)
}

func getStory(cmd *cobra.Command, args []string) error {
	var err error
	var resp *[]byte
	var absoluteOutput string
	var file *os.File
	var buf *bytes.Buffer

	if viper.GetString("output") != "" {
		absoluteOutput, err = filepath.Abs(viper.GetString("output"))
	}

	if err != nil {
		return err
	}

	if absoluteOutput != "" {
		file, err = os.OpenFile(absoluteOutput, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0622)
	}

	if err != nil {
		return err
	}

	resp, err = rest.Get(rest.GetIssues, http.DefaultClient, viper.GetViper())
	if err != nil {
		return err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(*resp, &m)
	if err != nil {
		return err
	}

	if viper.GetString("output") == "" {

	}

    // If the output flag is not provided print to stdout then quit.
    // Trying to keep the early return to prevent too many indentations.
	if absoluteOutput == "" {
		_, err = fmt.Println(extractIssues(m))
		return err
	}

    // Otherwise print to file
	buf, err = markdown.GenerateIssuesMarkdown(
		&markdown.Summary{
			Name:   filepath.Base(absoluteOutput),
			Issues: extractIssues(m),
		},
	)
	if err != nil {
		return err
	}

	_, err = file.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return err
}
