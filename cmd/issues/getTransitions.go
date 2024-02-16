package issues

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/cedrata/jira-helper/app/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getTransitionsCmd *cobra.Command

func init() {
	getTransitionsCmd = &cobra.Command{
		Use:   "get-transitions --issue-key <issue-key>",
		Short: "Get the issue transitions",
		Long: `Given an issue key or Id return the possible transitions from 
        the current one`,
		RunE: getTransitionsHandler,
	}

	// Path parameters
	// Append flags to command
	_ = getTransitionsCmd.
		Flags().
		String(
			"issue-key",
			"",
			"The ID or key of the issue.",
		)

	// Requied Flags
	_ = getTransitionsCmd.MarkFlagRequired("issue-key")

	// Bond flags viper
	_ = viper.BindPFlag(
		"issue-key",
		getTransitionsCmd.
			Flags().
			Lookup("issue-key"),
	)

	// Query parameters
	// Append flags to command
	_ = getTransitionsCmd.
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
	_ = getTransitionsCmd.
		Flags().
		String(
			"transition-id",
			"",
			"The ID of the transition.",
		)
	_ = getTransitionsCmd.
		Flags().
		Bool(
			"skip-remote-only-condition",
			false,
			`Whether transitions with the condition Hide From User Condition 
are included in the response.`,
		)
	_ = getTransitionsCmd.
		Flags().
		Bool(
			"include-unavailable-transitions",
			false,
			`Whether details of transitions that fail a condition are included 
in the response`,
		)
	_ = getTransitionsCmd.
		Flags().
		Bool(
			"sort-by-ops-bar-and-status",
			false,
			`Whether the transitions are sorted by ops-bar sequence value first
then category order (Todo, In Progress, Done) or only by ops-bar 
sequence value.`,
		)

	// Bond flags viper
	_ = viper.
		BindPFlag(
			"expand",
			getTransitionsCmd.
				Flags().
				Lookup("expand"),
		)
	_ = viper.
		BindPFlag(
			"transition-id",
			getTransitionsCmd.
				Flags().
				Lookup("transition-id"),
		)
	_ = viper.
		BindPFlag(
			"skip-remote-only-condition",
			getTransitionsCmd.
				Flags().
				Lookup("skip-remote-only-condition"),
		)
	_ = viper.
		BindPFlag(
			"include-unavailable-transitions",
			getTransitionsCmd.
				Flags().
				Lookup("include-unavailable-transitions"),
		)
}

func getTransitionsHandler(cmd *cobra.Command, args []string) error {
	var queryParameters = make(map[string]string)

	var successStatusCodes = []int{
		http.StatusOK,
	}
	var errorStatusCode = []int{
		http.StatusUnauthorized,
        http.StatusNotFound,
	}

	// add provided query parameters
	if viper.GetString("expand") != "" {
		queryParameters["expand"] = viper.GetString("expand")
	}

	if viper.GetString("transition-id") != "" {
		queryParameters["transition-id"] = viper.GetString("transition-id")
	}

	// Boolean properties are added anyway because they have a default value
	queryParameters["skip-remote-only-condition"] =
		viper.GetString("skip-remote-only-condition")

	queryParameters["include-unavailable-transitions"] =
		viper.GetString("include-unavailable-transitions")

	requestHelper := rest.NewRequestHelper(
		viper.GetString("host"),
		fmt.Sprintf(
			"rest/api/2/issue/%s/transitions",
			viper.GetString("issue-key"),
		),
		http.MethodGet,
		queryParameters,
		rest.GetHeadersWithBearer(
			viper.GetString("token"),
		),
		nil,
	)

	request, _ := requestHelper.BuildRequest()

	_, _ = http.DefaultClient.Do(request)

    // validate response and print value if any

	return nil
}
