package issuesearch

import (
	"github.com/cedrata/jira-helper/cmd/issuesearch/jql"
	"github.com/spf13/cobra"
)

var IssueSearchCmd *cobra.Command

func init() {
	IssueSearchCmd = &cobra.Command{
		Use: "issue-search",
	}

    IssueSearchCmd.AddCommand(jql.JqlCmd)
}
