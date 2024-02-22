package issues

import (
	"github.com/cedrata/jira-helper/cmd/issues/gettransition"
	"github.com/cedrata/jira-helper/cmd/issues/transitionissue"
	"github.com/spf13/cobra"
)

var IssuesCmd *cobra.Command

func init() {
	IssuesCmd = &cobra.Command{
		Use: "issues",
	}


	IssuesCmd.AddCommand(gettransition.GetTransitionsCmd)
    IssuesCmd.AddCommand(transitionissue.TransitionIssueCmd)
}
