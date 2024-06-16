package jql

import (
	"fmt"
	"net/http"

	"github.com/cedrata/jira-helper/pkg/config"
	httpHelpers "github.com/cedrata/jira-helper/pkg/helpers/http"
	"github.com/spf13/cobra"
)

var JqlCmd *cobra.Command

func init() {
	JqlCmd = &cobra.Command{
		Use:   "jql",
		Short: "Search for an issuea with a JQL query",
		Long:  "Search for an issuea with a JQL query",
		RunE:  jqlHandler,
	}

	// Body
	// Append flags to command
	_ = JqlCmd.Flags().String(
		"body",
		"",
		"A JSON object containing the search request.",
	)

	// Required flags
	_ = JqlCmd.MarkFlagRequired("body")
}

func jqlHandler(cmd *cobra.Command, args []string) error {
	var err error
	var body string

	// store body
	body, err = cmd.Flags().GetString("body")
	if err != nil {
		return err
	}

	requestHelper := httpHelpers.NewRequestHelper(
		config.ConfigData.Host,
		"rest/api/2/search",
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

	resopnse, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	message, err := httpHelpers.JSONHttpResponse(resopnse)
	if err != nil {
		return err
	}

	fmt.Println(message)

	return err
}
