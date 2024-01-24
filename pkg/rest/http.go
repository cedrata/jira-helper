package rest

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"text/template"

	"github.com/spf13/viper"
)

func Get(op Operation, client *http.Client, flags *viper.Viper) (*[]byte, error) {
	var payload = []byte("")
	var err error
	var url string
	var token string
	var req *http.Request
	var resp *http.Response

	url, err = operationSwitch(op, flags)
	if err != nil {
		return &payload, err
	}

	token = flags.GetString("token")
	if token == "" {
		return &payload, errors.New("token is missing")
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
		return &payload, err
	}

	return &payload, nil
}

func operationSwitch(op Operation, flags *viper.Viper) (string, error) {
	var builtUrl string
	var err error

	switch op {
	case GetIssues:
		var urlTemplate string
		var jiraUrl string

		jiraUrl = flags.GetString("host")
		if jiraUrl == "" {
			return "", errors.New("host is not provided, make sure \"host\" is provided with configuration file or flag")
		}

		content := urlSearch{JiraUrl: jiraUrl}

		urlTemplate, err = getUrlFromTemplate(TemplateUrlSearch, GetIssues, content)

		statements := []string{}
		fields := []string{"description", "status", "issueKey", "assignee", "summary"}
		if project := flags.GetString("project"); project != "" {
			statements = append(statements, fmt.Sprintf("project=%s", project))
		}

		if user := flags.GetString("user"); user != "" {
			statements = append(statements, fmt.Sprintf("assignee=%s", user))
		}

		if status := flags.GetString("status"); status != "" {
			statements = append(statements, "status="+url.PathEscape("\""+status+"\""))
		}

		if activeSprint := flags.GetBool("active-sprint"); activeSprint {
			statements = append(statements, "Sprint+in+openSprints()")
		}

		if types := flags.GetString("type"); types != "" {
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

	default:
		err = fmt.Errorf("unexpected operaion %s", op)
	}

	return builtUrl, err
}

func getUrlFromTemplate(t string, op Operation, content any) (string, error) {
	urlTemplate, err := template.
		New(fmt.Sprintf("operation-%s-url", op)).
		Parse(t)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	urlTemplate.Execute(buf, content)
	url := buf.String()
	return url, nil
}
