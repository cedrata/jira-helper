package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/cedrata/jira-helper/pkg/config"
	"github.com/cedrata/jira-helper/pkg/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	Execute()
}

var (
	// Used for flags.
	cfgFile     string
	userLicense string
	conf        *config.Config

	rootCmd = &cobra.Command{
		Use:   "jirahelper <subcommand>",
		Short: "An helper for using JIRA on CLI",
		Long:  `An helper for using JIRA on CLI`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")

	for i := range cmds {
		rootCmd.AddCommand(&cmds[i])
	}
}

func initConfig() {
	var err error

	fmt.Println("pippooo")
	if cfgFile != "" {
		// Use config file from the flag.
		conf, err = config.LoadLocalConfig(cfgFile, ".cobra.yaml")
		cobra.CheckErr(err)
		if err != nil {
			fmt.Printf("Error loading config file: %v\n", err)
		} else {
			fmt.Printf("Loaded Config: %+v\n", conf)
		}
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error finding home directory: %v\n", err)
		} else {
			fmt.Printf("Home Directory: %s\n", home)
		}
		cobra.CheckErr(err)

		conf, err = config.LoadLocalConfig(home, ".cobra.yaml")
		if err != nil {
			fmt.Printf("Error loading config file: %v\n", err)
		} else {
			fmt.Printf("Loaded Config: %+v\n", conf)
		}
		fmt.Println(conf)
		cobra.CheckErr(err)
	}
}

var cmds = []cobra.Command{
	{
		Use:   "issues <userId>",
		Short: "Get issues for a user",
		Long:  `This create a new agile documentation directory`,
		RunE:  getStory,
	},
}

func getStory(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("expected one element")
	}

	resp, err := rest.Get(rest.GetIssues, conf, http.DefaultClient)

	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}
