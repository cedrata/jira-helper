package rest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"text/template"

	"github.com/cedrata/jira-helper/pkg/config"
	"github.com/pkg/errors"
)

func Get(op Operation, c *config.Config, client *http.Client) (*[]byte, error) {
	var payload = []byte("")
	var err error
	var url string
	var req *http.Request
	var resp *http.Response

	url, err = operationSwitch(op, c)
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

func operationSwitch(op Operation, content any) (string, error) {
	var urlTemplate string
	var err error
	switch op {
	case GetIssues:
		urlTemplate, err = getUrlFromTemplate(TemplateUrlGetProjIssues, GetIssues, content)
	default:
		err = errors.Errorf("unexpected operaion %s", op)
	}

	if err != nil {
		return "", errors.WithStack(err)
	}

	return getUrlFromTemplate(urlTemplate, op, content)
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
