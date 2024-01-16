package jira

import "fmt"

type Issue struct {
	Key         string
	Assignee    string
	Description string
	Status      string
	Summary     string
}

type Issues []Issue

func (i Issue) String() string {
	return fmt.Sprintf(
		"key: %s\nsummary: %s\nassignee: %s\nstatus: %s\ndescription: %s",
		i.Key,
		i.Summary,
		i.Assignee,
		i.Status,
		i.Description,
	)
}

func (i Issues) String() string {
    var res string

    for k := range i {
        res = fmt.Sprintf("%s\n\n%s", res, i[k].String())
    }

    return res[2:]
}
