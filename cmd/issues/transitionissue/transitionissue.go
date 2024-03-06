package transitionissue

import (
	"fmt"
	"net/http"

	"github.com/cedrata/jira-helper/pkg/config"
	httpHelpers "github.com/cedrata/jira-helper/pkg/helpers/http"
	"github.com/spf13/cobra"
)

var TransitionIssueCmd *cobra.Command

func init() {
	TransitionIssueCmd = &cobra.Command{
		Use:   "transition-issue",
		Short: "Change status of the given issue",
		Long: `Performs an issue transition and, if the transition has a screen
, updates the fields from the transition screen`,
		RunE: transitionIssueHandler,
	}

	// Path parameters
	// Append flags to command
	_ = TransitionIssueCmd.
		Flags().
		String(
			"issue-key",
			"",
			"the ID or key of the issue.",
		)

	// Requied Flags
	_ = TransitionIssueCmd.MarkFlagRequired("issue-key")

	// Body
	// Append flags to command
	_ = TransitionIssueCmd.
		Flags().
		String("body", "", "The JSON body for the request")

	// Requied Flags
	_ = TransitionIssueCmd.MarkFlagRequired("body")
}

func transitionIssueHandler(cmd *cobra.Command, args []string) error {
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

	requestHelper := httpHelpers.NewRequestHelper(
		config.ConfigData.Host,
		fmt.Sprintf(
			"rest/api/2/issue/%s/transitions",
			issueKey,
		),
		http.MethodPost,
		make(map[string]string),
		httpHelpers.PostHeadersWithBearer(
			config.ConfigData.Token,
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

	message, err := httpHelpers.JSONHttpReponse(response)
	if err != nil {
		return err
	}

	fmt.Println(message)

	return nil
}
