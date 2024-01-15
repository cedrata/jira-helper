package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cedrata/jira-helper/pkg/jira"
	"github.com/cedrata/jira-helper/pkg/markdown"
	"github.com/cedrata/jira-helper/pkg/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var testsCmd *cobra.Command

func init() {
	testsCmd = &cobra.Command{
		Use:   "tests [flags] <file>",
		Short: "Create a test table with all tests",
		Long:  `This create a test table in markdown with all tests`,
		RunE:  writeTestList,
	}

	testsCmd.Flags().String("type", "test", "select the issue type")
	viper.BindPFlag("type", issuesCmd.Flags().Lookup("type"))

	rootCmd.AddCommand(testsCmd)
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
