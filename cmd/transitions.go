package cmd

import (
	"fmt"
	"net/http"

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
	transitionsCmd.Flags().String("to-transition", "", "if provided try to set the issue to the given transition id")

    _ = viper.BindPFlag("key", transitionsCmd.Flags().Lookup("key"))
    _ = viper.BindPFlag("to-transition", transitionsCmd.Flags().Lookup("to-transition"))

	rootCmd.AddCommand(transitionsCmd)
}

func handleTransitions(cmd *cobra.Command, args []string) error {
	// If to-transition == "" then GET
	// Else POST
	fmt.Printf("returning possible transitions for issue having id %s\n", viper.GetString("key"))
	resp, err := rest.Get(rest.GetTransitions, http.DefaultClient, viper.GetViper())
	if err != nil {
		fmt.Printf("%s", *resp)
	}
	return err
}
