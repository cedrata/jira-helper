package gettransition

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cedrata/jira-helper/pkg/config"
	httpHelpers "github.com/cedrata/jira-helper/pkg/helpers/http"
	"github.com/spf13/cobra"
)

var GetTransitionsCmd *cobra.Command

func init() {
	GetTransitionsCmd = &cobra.Command{
		Use:   "get-transitions",
		Short: "Get the issue transitions",
		Long: `Given an issue key or Id return the possible transitions from 
        the current one`,
		RunE: getTransitionsHandler,
	}

	// Path parameters
	// Append flags to command
	_ = GetTransitionsCmd.
		Flags().
		String(
			"issue-key",
			"",
			"The ID or key of the issue.",
		)

	// Requied Flags
	_ = GetTransitionsCmd.MarkFlagRequired("issue-key")

	// Query parameters
	// Append flags to command
	_ = GetTransitionsCmd.
		Flags().
		String(
			"expand",
			"",
			`Use expand to include additional information about transitions in 
the response. This parameter accepts transitions.fields, which 
returns information about the fields in the transition screen for 
each transition. Fields hidden from the screen are not returned. 
Use this information to populate the fields and update fields in 
Transition issue.`,
		)
	_ = GetTransitionsCmd.
		Flags().
		String(
			"transition-id",
			"",
			"The ID of the transition.",
		)
	_ = GetTransitionsCmd.
		Flags().
		Bool(
			"skip-remote-only-condition",
			false,
			`Whether transitions with the condition Hide From User Condition 
are included in the response.`,
		)
	_ = GetTransitionsCmd.
		Flags().
		Bool(
			"include-unavailable-transitions",
			false,
			`Whether details of transitions that fail a condition are included 
in the response`,
		)
	_ = GetTransitionsCmd.
		Flags().
		Bool(
			"sort-by-ops-bar-and-status",
			false,
			`Whether the transitions are sorted by ops-bar sequence value first
then category order (Todo, In Progress, Done) or only by ops-bar 
sequence value.`,
		)
}

func getTransitionsHandler(cmd *cobra.Command, args []string) error {
	var queryParameters = make(map[string]string)
	var err error
	var issueKey string
	var expand string
	var transitionId string
	var skipRemoteOnlyCondition bool
	var includeUnavailableTransitions bool

	// store path parameters
	issueKey, err = cmd.Flags().GetString("issue-key")
	if err != nil {
		return err
	}

	// add provided query parameters
	expand, err = cmd.Flags().GetString("expand")
	if err != nil {
		return err
	}

	if expand != "" {
		queryParameters["expand"] = expand
	}

	transitionId, err = cmd.Flags().GetString("transition-id")
	if err != nil {
		return err
	}

	if transitionId != "" {
		queryParameters["transition-id"] = transitionId
	}

	// boolean properties are added anyway because they have a default value
	skipRemoteOnlyCondition, err =
		cmd.Flags().GetBool("skip-remote-only-condition")
	if err != nil {
		return err
	}
	queryParameters["skip-remote-only-condition"] =
		strconv.FormatBool(skipRemoteOnlyCondition)

	includeUnavailableTransitions, err =
		cmd.Flags().GetBool("include-unavailable-transitions")
	if err != nil {
		return err
	}
	queryParameters["include-unavailable-transitions"] =
		strconv.FormatBool(includeUnavailableTransitions)

	requestHelper := httpHelpers.NewRequestHelper(
		config.ConfigData.Host,
		fmt.Sprintf(
			"rest/api/2/issue/%s/transitions",
			issueKey,
		),
		http.MethodGet,
		queryParameters,
		httpHelpers.GetHeadersWithBearer(
			config.ConfigData.Token,
		),
		nil,
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
