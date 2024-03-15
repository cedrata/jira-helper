package main

import (
	"github.com/cedrata/jira-helper/cmd"
	"github.com/spf13/cobra"
)

func main() {
	if err := cmd.Execute(); err != nil {
		cobra.CheckErr(err)
	}
}
