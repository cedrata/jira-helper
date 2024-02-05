package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/cedrata/jira-helper/app/markdown"
	"github.com/cedrata/jira-helper/app/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var newDocCmd *cobra.Command

func init() {
	newDocCmd = &cobra.Command{
		Use:   "doc [flags] <dir> <story id>",
		Short: "Create a new doc directory",
		Long: `This create a new agile documentation directory with a 
        presentation and a documentation markdown file`,
		RunE: writeStoryTemplate,
	}

	newDocCmd.Flags().StringP("user", "u", "AF82260", "user name to filter issues for")
	newDocCmd.Flags().BoolP("active-sprint", "a", true, "select the issues only in active sprints")

	_ = viper.BindPFlag("user", newDocCmd.Flags().Lookup("user"))
	_ = viper.BindPFlag("active-sprint", newDocCmd.Flags().Lookup("active-sprint"))

	rootCmd.AddCommand(newDocCmd)
}

func writeStoryTemplate(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("expected  2 args found %d", len(args))
	}

	dir := args[0]
	id := args[1]

	resp, err := rest.Get(rest.GetIssues, http.DefaultClient, viper.GetViper())
	if err != nil {
		return err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(*resp, &m)
	if err != nil {
		return err
	}

	issues := extractIssuesMap(m)
	story, ok := issues[id]
	if !ok {
		return fmt.Errorf("error key %s not found", args[0])
	}

	err = os.Mkdir(dir, 0o777)
	if err != nil {
		return err
	}

	err = markdown.WriteStub(story, path.Join(dir, markdown.DocFileSrc),
		markdown.DocTemplate)
	if err != nil {
		return err
	}

	err = markdown.WriteStub(story, path.Join(dir, markdown.PresFileSrc),
		markdown.PresTemplate)
	if err != nil {
		return err
	}

	fmt.Printf("created doc dir %s\n", args[0])
	return nil
}
