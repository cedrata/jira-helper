package rest

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

func Get(op Operation, client *http.Client, v *viper.Viper) (*[]byte, error) {
	var payload = []byte("")
	var err error
	var url string
	var token string
	var req *http.Request
	var resp *http.Response

	url, err = operationSwitch(op, v)
	if err != nil {
		return &payload, err
	}

	token = v.GetString("token")
	if token == "" {
		return &payload, errors.New("token is missing")
	}

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &payload, err
	}

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

// func Post(op Operation, client *http.Client, flags *viper.Viper, body *[]byte) (*[]byte, error) {
//     var payload = []byte("")
//     var err error
//     var url string
// }

func operationSwitch(op Operation, v *viper.Viper) (string, error) {
	var builtUrl string
	var err error

	switch op {
	case GetIssues:
		var urlTemplate string
		var host string

		host = v.GetString("host")
		if host == "" {
			return "", errors.New("host is not provided, make sure \"host\" is provided with configuration file or flag")
		}

		urlTemplate = fmt.Sprintf("https://%s/rest/api/2/search", host)
		statements := []string{}
		fields := []string{"description", "status", "issueKey", "assignee", "summary"}
		if project := v.GetString("project"); project != "" {
			statements = append(statements, fmt.Sprintf("project=%s", project))
		}

		if user := v.GetString("user"); user != "" {
			statements = append(statements, fmt.Sprintf("assignee=%s", user))
		}

		if status := v.GetString("status"); status != "" {
			statements = append(statements, "status="+url.PathEscape("\""+status+"\""))
		}

		if activeSprint := v.GetBool("active-sprint"); activeSprint {
			statements = append(statements, "Sprint+in+openSprints()")
		}

		if types := v.GetString("type"); types != "" {
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

	case GetTransitions:
		builtUrl, err = transitionsUrl(v)

	case PostTransitions:
		builtUrl, err = transitionsUrl(v)

	default:
		err = fmt.Errorf("unexpected operaion %s", op)
	}

	return builtUrl, err
}

func transitionsUrl(v *viper.Viper) (string, error) {
	var host string
	var issueKey string

	host = v.GetString("host")
	if host == "" {
		return "", errors.New("host is not provided, make sure \"host\" is provided with configuration file or flag")
	}

	issueKey = v.GetString("key")
	if issueKey == "" {
		return "", errors.New("key is not provided, make sure \"key\" is provided with configuration file or flag")
	}

	return fmt.Sprintf("https://%s/rest/api/2/issue/%s/transitions", host, issueKey), nil
}
