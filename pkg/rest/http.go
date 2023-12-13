package rest

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	JSONContentType = "application/json"
)

type RestClient struct {
	Token   string
	Project string
	Url     string
}

func (client *RestClient) GetAllIssues() error {
	url := client.Url + "/rest/api/2/search?jql=project=" + client.Project + "+order+by+duedate&fields=id,key"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	req.Header.Set("Authorization", "Bearer " + client.Token)
	req.Header.Set("Content-Type", JSONContentType)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}
	
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status ok found %s", http.StatusText(res.StatusCode))
	}

	data, err  := io.ReadAll(res.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	fmt.Printf("data %s", string(data))

	return nil
}
