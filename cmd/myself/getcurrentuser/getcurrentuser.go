package getcurrentuser

import (
	"fmt"
	"net/http"

	"github.com/cedrata/jira-helper/pkg/config"
	httpHelpers "github.com/cedrata/jira-helper/pkg/helpers/http"
	"github.com/spf13/cobra"
)

var GetCurrentUserCmd *cobra.Command

func init() {
	GetCurrentUserCmd = &cobra.Command{
		Use:   "get-current-user",
		Short: "Returns details for the current user",
		Long:  "Returns details for the current user",
		RunE:  getCurrentUserHandler,
	}

	// Query parameters
	_ = GetCurrentUserCmd.
		Flags().
		String(
			"expand",
			"",
			`Use expand to include additional information about user in the 
response. This parameter accepts a comma-separated list. 
Expand options include: 
* groups Returns all groups, including nested groups, 
the user belongs to 
* applicationRoles Returns the application roles the user 
is assigned to.
`)
}

func getCurrentUserHandler(cmd *cobra.Command, flags []string) error {
	var queryParameters = make(map[string]string)
	var err error
	var expand string

	// add provided query parameters
	expand, err = cmd.Flags().GetString("expand")
	if err != nil {
		return err
	}

	if expand != "" {
		queryParameters["expand"] = expand
	}

	requestHelper := httpHelpers.NewRequestHelper(
		config.ConfigData.Host,
		"/rest/api/2/myself",
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

	message, err := httpHelpers.JSONHttpResponse(response)
	if err != nil {
		return err
	}

	fmt.Println(message)

	return nil
}
