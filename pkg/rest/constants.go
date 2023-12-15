package rest

const (
	JSONContentType = "application/json"
	// Host contains protocol and hostname
    TemplateUrlGetProjIssues = "{{.Protocol}}://{{.Host}}/rest/api/2/search?jql=project={{.ProjectId}}+order+by+duedate&fields=id,key"
)

const (
	GetIssues Operation = "get-issues"
)
