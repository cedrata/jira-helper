package cmd

import (
	"encoding/json"
	"errors"
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
	_ = viper.BindPFlag("to", transitionsCmd.Flags().Lookup("to-transition"))

	rootCmd.AddCommand(transitionsCmd)
}

func handleTransitions(cmd *cobra.Command, args []string) error {
	var err error
	var resp *[]byte

	if viper.GetString("to") == "" {
		fmt.Printf("returning possible transitions for issue having id %s\n", viper.GetString("key"))
		resp, err = rest.Get(rest.GetTransitions, http.DefaultClient, viper.GetViper())
	} else {
		return errors.New("Operation PostTransitions not implemented yet")
	}

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

	return nil
}
