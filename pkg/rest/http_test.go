package rest

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllIssues(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", JSONContentType)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"pippo": "pippo"}`))

		t.Logf("Requested path: %s", r.URL)

		// // Check the request method and path
		// assert.Equal(t, http.MethodGet, r.Method)
		// assert.Equal(t, "/rest/api/2/search?jql=project=your_project+order+by+duedate&fields=id,key", r.URL.Path)

		// // Respond with a sample JSON response
		// w.Header().Set("Content-Type", JSONContentType)
		// w.WriteHeader(http.StatusOK)
		// fmt.Fprint(w, `{"key": "value"}`)
	}))
	defer server.Close()

	parsedUrl, _ := url.Parse(server.URL)
	payload, err := Get(
		JiraConfig{
			Token:     "token",
			Protocol:  "http",
			Host:      parsedUrl.Host,
			Operation: GetIssues,
			ProjectId: "INAEDM",
		},
	)

	t.Logf("payload: %s", payload)

	assert.Nil(t, err)
	assert.NotEmpty(t, payload)
}
