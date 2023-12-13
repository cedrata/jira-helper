package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllIssues(t *testing.T) {
	client := RestClient{
		Url:     "https://jira.springlab.enel.com",
		Project: "INAEDM",
		Token:   "NTY2MzgyODQ4NjgwOn/Kr/wOeoysnAyb5CsPczpzgPup",
	}

	// Create a test HTTP server to handle the request
	/*
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check the request method and path
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "/rest/api/2/search?jql=project=your_project+order+by+duedate&fields=id,key", r.URL.Path)

			// Respond with a sample JSON response
			w.Header().Set("Content-Type", JSONContentType)
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"key": "value"}`)
		}))
		defer server.Close()
		client.Url = server.URL
	*/

	err := client.GetAllIssues()
	assert.NoError(t, err)
}
