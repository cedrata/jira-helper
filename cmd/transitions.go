package cmd

import (
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

	viper.BindPFlag("key", issuesCmd.Flags().Lookup("key"))
	viper.BindPFlag("to-transition", issuesCmd.Flags().Lookup("to-transition"))
}

func handleTransitions(cmd *cobra.Command, args []string) error {
    // If to-transition == "" then GET
    // Else POST
	return nil
}
