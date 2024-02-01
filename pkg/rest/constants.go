package rest

type Operation string

const (
	JSONContentType = "application/json"

	// Available operations
	GetIssues       Operation = "get-issues"
	GetTransitions  Operation = "get-transitions"
	PostTransitions Operation = "post-transitions"
)
