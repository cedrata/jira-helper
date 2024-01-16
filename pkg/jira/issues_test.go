package jira

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintIssue(t *testing.T) {
	type Tester struct {
		input    Issue
		expected string
		name     string
	}

	var tests = []Tester{
		{
			input: Issue{
				Key:         "1",
				Summary:     "2",
				Assignee:    "3",
				Status:      "4",
				Description: "5",
			},
			expected: "key: 1\nsummary: 2\nassignee: 3\nstatus: 4\ndescription: 5",
			name:     "Test Issue String method OK",
		},
	}

	for test := range tests {
		current := tests[test]
		t.Run(current.name, (func(t *testing.T) {
			t.Parallel()

			actual := fmt.Sprint(current.input)
			t.Logf("Input %v\nresult in %s", current.input, actual)
			assert.Equal(t, current.expected, actual)
		}))
	}

}

func TestPrintIssues(t *testing.T) {
	type Tester struct {
		input    Issues
		expected string
		name     string
	}

	var tests = []Tester{
		{
			input: []Issue{
				{
					Key:         "1",
					Summary:     "2",
					Assignee:    "3",
					Status:      "4",
					Description: "5",
				},
				{
					Key:         "6",
					Summary:     "7",
					Assignee:    "8",
					Status:      "9",
					Description: "10",
				},
			},
			expected: `key: 1
summary: 2
assignee: 3
status: 4
description: 5

key: 6
summary: 7
assignee: 8
status: 9
description: 10`,
			name: "TestIssuesStringmethodOK",
		},
	}

	for test := range tests {
		current := tests[test]
		t.Run(current.name, (func(t *testing.T) {
			t.Parallel()

			actual := fmt.Sprint(current.input)
			t.Logf("Input %v\nresult in %s", current.input, actual)
			assert.Equal(t, current.expected, actual)
		}))
	}

}
