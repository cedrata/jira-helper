package issues

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getTransitionsCmd *cobra.Command

func init() {
	getTransitionsCmd = &cobra.Command{
		Use:   "get-transitions --issue-key <issue-key>",
		Short: "Get the issue transitions",
		Long:  "Given an issue key or Id return the possible transitions from the current one",
		RunE:  getTransitionsHandler,
	}

    // @TODO: Find less repetitive way to add flags like a function that takes pointer of command and viper then adds everything conditionally given some more arguments
	// Path parameters
	_ = getTransitionsCmd.Flags().String("issue-key", "", "key or id of the issue to return the transitions for")

	_ = getTransitionsCmd.MarkFlagRequired("issue-key")

	_ = viper.BindPFlag("issue-key", getTransitionsCmd.Flags().Lookup("issue-key"))

	// Query parameters
}

func getTransitionsHandler(cmd *cobra.Command, args []string) error {
	return nil
}
