package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cedrata/jira-helper/pkg/jira"
	"github.com/cedrata/jira-helper/pkg/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var transitionsCmd *cobra.Command

func init() {
	transitionsCmd = &cobra.Command{
		Use:   "transitions [flags]",
		Short: "Get and set the transitions of for a given issue",
		Long: `Get the transitions for an issue given it's key, if a new 
        a new transition is defined tries to transition to it`,
		RunE: handleTransitions,
	}

	transitionsCmd.Flags().String("key", "", "issue key to get or set transitions for")
	transitionsCmd.Flags().String("to", "", "if provided try to set the issue to the given transition id")

	_ = viper.BindPFlag("key", transitionsCmd.Flags().Lookup("key"))
	_ = viper.BindPFlag("to", transitionsCmd.Flags().Lookup("to"))

	rootCmd.AddCommand(transitionsCmd)
}

func handleTransitions(cmd *cobra.Command, args []string) error {
	var err error
	var resp *[]byte

	if viper.GetString("to") == "" {
		fmt.Printf("returning possible transitions for issue having id %s\n", viper.GetString("key"))
		resp, err = rest.Get(rest.GetTransitions, http.DefaultClient, viper.GetViper())

		if err != nil {
			return err
		}
		var transitions jira.GetTransitionResponse
		if err = json.Unmarshal(*resp, &transitions); err != nil {
			return err
		}

		for _, transition := range transitions.Transitions {
			fmt.Printf("id: %s\nname: %s\n\n", transition.Id, transition.Name)
		}
	} else {
		bodyObj := &jira.PostTransitionBody{Transition: &jira.IssueTransition{Id: viper.GetString("to")}}
		body, _ := json.Marshal(bodyObj)
		_, err = rest.Post(rest.PostTransitions, http.DefaultClient, viper.GetViper(), &body)

		if err != nil {
			fmt.Println("Some error occured while executing the transition, visit the following link to see the possible causes: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-issueidorkey-transitions-post-response")
		} else {
            fmt.Println("Succesful state transition")
        }
	}

	return nil
}
