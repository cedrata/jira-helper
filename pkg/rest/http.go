package rest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

func Get(op Operation, client *http.Client, flags *pflag.FlagSet) (*[]byte, error) {
	var payload = []byte("")
	var err error
	var url string
	var token string
	var req *http.Request
	var resp *http.Response

	url, err = operationSwitch(op, flags)
	if err != nil {
		return &payload, errors.WithStack(err)
	}

	token, err = flags.GetString("token")
	if err != nil {
		return &payload, err
	}

	req, err = http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", JSONContentType)
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

func operationSwitch(op Operation, flags *pflag.FlagSet) (string, error) {
	var builtUrl string
	var err error

	switch op {
	case GetIssues:
		var urlTemplate string
		var jiraUrl string

		jiraUrl, err = flags.GetString("host")
		if err != nil {
			break
		}

		content := urlSearch{JiraUrl: jiraUrl}

		urlTemplate, err = getUrlFromTemplate(TemplateUrlSearch, GetIssues, content)

		statements := []string{}
		fields := []string{"description", "status", "issueKey", "assignee", "summary"}
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

		if types, err := flags.GetString("type"); err == nil && types != "" {
			statements = append(statements, "issueType="+types)
		}

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
				"maxResults=70",
			}, "&",
		)

		builtUrl = strings.Join(
			[]string{
				urlTemplate,
				query,
			}, "?",
		)

		fmt.Printf("url %s\n", builtUrl)

	default:
		err = errors.Errorf("unexpected operaion %s", op)
	}

	return builtUrl, errors.WithStack(err)
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
