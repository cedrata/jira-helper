package markdown

import (
	"testing"

	"github.com/cedrata/jira-helper/pkg/jira"
	"github.com/stretchr/testify/assert"
)

func TestGenerateIssueString(t *testing.T) {
	type Tester struct {
		input    Summary
		expected string
		name     string
	}

	var tests = []Tester{
		{
			input: Summary{
				Issues: jira.Issues{
					{
						Key:         "key-1",
						Summary:     "summary number one",
						Assignee:    "me",
						Status:      "to do",
						Description: "A longer and more hexaustive description for first issue",
					},
				},
				Name: "first",
			},

			expected: `# first

## key-1 - summary number one

### description
A longer and more hexaustive description for first issue

### conclusion/tests


`,
			name: "SingleIssueSummaryString",
		},
		{
			input: Summary{
				Issues: jira.Issues{
					{
						Key:         "key-1",
						Summary:     "summary number one",
						Assignee:    "me",
						Status:      "to do",
						Description: "A longer and more hexaustive description for first issue",
					},
					{
						Key:         "key-2",
						Summary:     "summary number two",
						Assignee:    "me",
						Status:      "doing",
						Description: "A longer and more hexaustive description for second issue",
					},
				},
				Name: "second",
			},

			expected: `# second

## key-1 - summary number one

### description
A longer and more hexaustive description for first issue

### conclusion/tests


## key-2 - summary number two

### description
A longer and more hexaustive description for second issue

### conclusion/tests


`,
			name: "MultipleIssueSummaryString",
		},
	}

	for test := range tests {
		current := tests[test]
		t.Run(current.name, func(t *testing.T) {
			actual, err := GenerateIssuesString(&current.input)
			assert.NoError(t, err)
			assert.Equal(t, current.expected, actual.String())
		})
	}
}
