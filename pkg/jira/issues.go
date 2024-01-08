package jira

import (
	"fmt"
)

type Issue struct {
	Key         string
	Assignee    string
	Description string
	Status      string
	Summary     string
}

func (i Issue) String() string {
	return fmt.Sprintf(
		"\nkey: %s\nsummary:%s\nassignee: %s\nstatus: %s\ndescription: %s\n",
		i.Key,
		i.Summary,
		i.Assignee,
		i.Status,
		i.Description,
	)
}
