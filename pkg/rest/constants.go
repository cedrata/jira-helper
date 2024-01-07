package rest

const (
	JSONContentType          = "application/json"
	TemplateUrlSearch = "https://{{.JiraUrl}}/rest/api/2/search" //?jql=project={{.Project}}+order+by+duedate&fields=id,key"
)

const (
	GetIssues Operation = "get-issues"
)
