package rest

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildRequest(t *testing.T) {
	type tester struct {
		helper            RequestHelper
		expectedUrlString string
		name              string
	}

	tests := []tester{
		{
			RequestHelper{
				"sample.host.com",
				"resource/foo",
				"https",
				"",
				map[string]string{"key": "value"},
				map[string]string{},
				[]byte(""),
			},
			"https://sample.host.com/resource/foo?key=value",
			"good-url-request-no-trailing-slash",
		},
		{
			RequestHelper{
				"sample.host.com/",
				"/resource/foo/",
				"https",
				"",
				map[string]string{"key": "value"},
				map[string]string{},
				[]byte(""),
			},
			"https://sample.host.com/resource/foo/?key=value",
			"good-url-request-end-slash-host-beginning-slash-resource",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := test.helper.BuildRequest()

			assert.NoError(t, err)
			assert.Equal(t, test.expectedUrlString, request.URL.String())
		})
	}
}

func TestJSONPrettyResponse(t *testing.T) {
	type result struct {
		body string
		err  error
	}

	type tester struct {
		response *http.Response
		expected result
		name     string
	}

	tests := []tester{
		{
			response: &http.Response{
				StatusCode: 204,
				Body:       io.NopCloser(bytes.NewBufferString("")),
			},
			expected: result{
				body: `{
  "statusCode": 204,
  "body": {}
}`,
				err: nil,
			},
			name: "no content body",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, err := JSONHttpReponse(test.response)

			assert.Equal(t, test.expected.err, err)
			assert.Equal(t, test.expected.body, body)
		})
	}
}
