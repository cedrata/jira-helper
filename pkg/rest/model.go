package rest

type Operation string

type JiraConfig struct {
	// Jira instance host name with protocol e.g.:
	Host string
    // http or https
    Protocol string
	// Jira access token
	Token string
	// Jira project id
	ProjectId string
	Operation
}

// type GetIssues struct {
// 	JiraBase
// 	ProjectId string
// }
