package rest

import (
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
