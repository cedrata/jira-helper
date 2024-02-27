package assignissue

import (
	"fmt"
	"net/http"

	"github.com/cedrata/jira-helper/pkg/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var AssignIssueCmd *cobra.Command

func init() {
	AssignIssueCmd = &cobra.Command{
		Use:   "assign-issue",
		Short: "Assign an issue to a given user",
		Long: `Assigns an issue to a user. Use this operation when the calling 
user does not have the Edit Issues permission but has the Assign issue 
permission for the project that the issue is in.`,
		RunE: assignIssueHandler,
	}

	// Path parameters
	// Append flags to command
	_ = AssignIssueCmd.
		Flags().
		String(
			"issue-key",
			"",
			"The ID or key of the issue.",
		)

	// Requied Flags
	_ = AssignIssueCmd.MarkFlagRequired("issue-key")

	// Body
	// Append flags to command
	_ = AssignIssueCmd.
		Flags().
		String("body", "", "The JSON body for the request")

	// Requied Flags
	_ = AssignIssueCmd.MarkFlagRequired("body")
}

func assignIssueHandler(cmd *cobra.Command, args []string) error {
	var err error
	var issueKey string
	var body string

	// store path parameters
	issueKey, err = cmd.Flags().GetString("issue-key")
	if err != nil {
		return err
	}

	// store body
	body, err = cmd.Flags().GetString("body")
	if err != nil {
		return err
	}

	requestHelper := rest.NewRequestHelper(
		viper.GetString("host"),
		fmt.Sprintf(
			"rest/api/2/issue/%s/assignee",
			issueKey,
		),
		http.MethodPut,
		make(map[string]string),
		rest.PutHeadersWithBearer(
			viper.GetString("token"),
		),
		[]byte(body),
	)

	request, err := requestHelper.BuildRequest()
	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	message, err := rest.JSONHttpReponse(response)
	if err != nil {
		return err
	}

	fmt.Println(message)

	return nil
}
