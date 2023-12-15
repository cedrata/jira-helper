package rest

const (
	JSONContentType = "application/json"
    TemplateUrlGetProjIssues = "{{.Protocol}}://{{.Host}}/rest/api/2/search?jql=project={{.ProjectId}}+order+by+duedate&fields=id,key"
)

const (
	GetIssues Operation = "get-issues"
)
