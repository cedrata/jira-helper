package rest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"text/template"

	"github.com/cedrata/jira-helper/pkg/config"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

func Get(op Operation, c *config.Config, client *http.Client, flags *pflag.FlagSet) (*[]byte, error) {
	var payload = []byte("")
	var err error
	var url string
	var req *http.Request
	var resp *http.Response

	url, err = operationSwitch(op, c, flags)
	if err != nil {
		return &payload, errors.WithStack(err)
	}

	req, err = http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", JSONContentType)
	req.Header.Set("User-Agent", "Go_JiraHelper/1.0")
	resp, err = client.Do(req)
	if err != nil {
		return &payload, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &payload, fmt.Errorf("expected status %s found %s",
			http.StatusText(http.StatusOK),
			http.StatusText(resp.StatusCode),
		)
	}

	payload, err = io.ReadAll(resp.Body)
	if err != nil {
		return &payload, errors.WithStack(err)
	}

	return &payload, nil
}

func operationSwitch(op Operation, content any, flags *pflag.FlagSet) (string, error) {
	var builtUrl string
	var err error

	switch op {
	case GetIssues:
		// URL example:
		// "https://jira.springlab.enel.com/rest/api/2/search?jql=project=${project}+AND+Sprint+in+openSprints()+AND+assignee=${jira_user}&fields=description,status"
		var urlTemplate string
		urlTemplate, err = getUrlFromTemplate(TemplateUrlSearch, GetIssues, content)

		// Generation of the JQL query
		// prefix := "jql="
		// suffix := ""
		statements := []string{}
		fields := []string{"description", "status", "issueKey", "assignee"}
		if project, err := flags.GetString("project"); err == nil && project != "" {
			statements = append(statements, fmt.Sprintf("project=%s", project))
		}

		if user, err := flags.GetString("user"); err == nil && user != "" {
			statements = append(statements, fmt.Sprintf("assignee=%s", user))
		}

		if status, err := flags.GetString("status"); err == nil && status != "" {
			statements = append(statements, "status="+url.PathEscape("\""+status+"\""))
		}

		if activeSprint, err := flags.GetBool("active-sprint"); err == nil && activeSprint == true {
			statements = append(statements, fmt.Sprintf("Sprint+in+openSprints()"))
		}

		fmt.Printf("statements: %s\n", statements)

		query := strings.Join(
			[]string{
				strings.Join(
					[]string{
						"jql",
						strings.Join(statements, "+AND+"),
					}, "=",
				),
				strings.Join(
					[]string{
						"fields",
						strings.Join(fields, ","),
					}, "=",
				),
			}, "&",
		)

		builtUrl = strings.Join(
			[]string{
				urlTemplate,
				query,
			},
			"?",
		)

		fmt.Println(builtUrl)
		return "", errors.Errorf("hello")

	default:
		err = errors.Errorf("unexpected operaion %s", op)
	}

	if err != nil {
		return "", errors.WithStack(err)
	}

	return getUrlFromTemplate(builtUrl, op, content)
}

func getUrlFromTemplate(t string, op Operation, content any) (string, error) {
	urlTemplate, err := template.
		New(fmt.Sprintf("operation-%s-url", op)).
		Parse(t)
	if err != nil {
		return "", errors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	urlTemplate.Execute(buf, content)
	url := buf.String()
	return url, nil
}
