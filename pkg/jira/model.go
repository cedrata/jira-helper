package jira

import "fmt"

type TestList struct {
	Tests []Issue
}

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

// Response from GET https://{host}/rest/api/3/issue/{issueIdOrKey}/transitions
type GetTransitionResponse struct {
	// Expand options that include additional transitions details in the response.
	Expand string `json:"expand"`

	// List of issue transitions.
	Transitions []IssueTransition `json:"transitions"`
}

//go:generate go run ../../main/generator -struct IssueTransition -unmarshal
type IssueTransition struct {
	// Expand options that include additional transition details in the response.
	Expand string `json:"expand"`

	// Details of the fields associated with the issue transition screen.
	// Use this information to populate fields and update in a transition request.
	Fields interface{} `json:"fields"`

	// Whether there is a screen associated with the issue transition.
	HasScreen bool `json:"hasScreen"`

	// The ID of the issue transition.
	// Required when specifying a transition to undertake.
	Id string `json:"id"`

	// Whether the transition is available to be performed.
	IsAvailable bool `json:"isAvailable"`

	// Whether the issue has to meet criteria before the issue transition is applied.
	IsConditional bool `json:"isConditional"`

	// Whether the issue transition is global, that is,
	// the transition is applied to issues regardless of their status.
	IsGlobal bool `json:"isGlobal"`

	// Whether this is the initial issue transition for the workflow.
	IsInitial bool `json:"isInitial"`

	Looped bool `json:"looped"`

	// The name of the issue transition
	Name string `json:"name"`

	// Details of the issue status after the transition.
	To StatusDetails `json:"to"`

	// Extra properties of any type may be provided to this object.
	AdditionalProperties map[string]interface{} `json:"-"`
}

//go:generate go run ../../main/generator -struct StatusDetails  -unmarshal
type StatusDetails struct {
	// The description of the status.
	Description string `json:"description"`

	// The URL of the icon used to represent the status.
	IconUrl string `json:"iconUrl"`

	// The ID of the status.
	Id string `json:"id"`

	// The name of the status.
	Name string `json:"name"`

	// The scope of the status.
	Scope Scope `json:"scope"`

	// The URL of the status.
	Self string `json:"self"`

	// The category assigned to the status.
	StatusCategory StatusCategory `json:"statusCategory"`

	// Extra properties of any type may be provided to this object.
	AdditionalProperties map[string]interface{} `json:"-"`
}

//go:generate go run ../../main/generator -struct StatusCategory -unmarshal
type StatusCategory struct {
    // The name of the color used to represent the status category.
    ColorName string `json:"colorName"`

    // The ID of the status category.
    // Format: int64
    Id int64 `json:"id"`

    // The key of the status category.
    Key string `json:"key"`

    // The name of the status category.
    Name string `json:"name"`

    // The URL of the status category.
    Self string `json:"self"`

	// Extra properties of any type may be provided to this object.
	AdditionalProperties map[string]interface{} `json:"-"`
}

//go:generate go run ../../main/generator -struct Scope -unmarshal
type Scope struct {
	// The project the item has scope in.
	Project ProjectDetails `json:"project"`

	// The type of scope.
	// Valid values: PROJECT, TEMPLATE
	Type string `json:"type"`

	// Extra properties of any type may be provided to this object.
	AdditionalProperties map[string]interface{} `json:"-"`
}

type ProjectDetails struct {
	// The URLs of the project's avatars.
	AvatarUrls AvatarUrlsBean `json:"avatarUrls"`

	// The ID of the project.
	Id string `json:"id"`

	// The key of the project.
	Key string `json:"key"`

	// The name of the project.
	Name string `json:"name"`

	// The category the project belongs to.
	ProjectCategory UpdatedProjectCategory `json:"projectCategory"`

	// The project type of the project.
	// Valid values: software, service_desk, business
	ProjectTypeKey string `json:"projectTypeKey"`

	// The URL of the project details.
	Self string `json:"self"`

	// Whether or not the project is simplified.
	Simplified bool `json:"simplified"`
}

type AvatarUrlsBean struct {
	// The URL of the item's 16x16 pixel avatar.
	Url16x16 string `json:"16x16"`

	// The URL of the item's 24x24 pixel avatar.
	Url24x24 string `json:"24x24"`

	// The URL of the item's 32x32 pixel avatar.
	Url32x32 string `json:"32x32"`

	// The URL of the item's 48x48 pixel avatar.
	Url48x48 string `json:"48x48"`
}

type UpdatedProjectCategory struct {
	// The name of the project category.
	Description string `json:"description"`

	// The ID of the project category.
	Id string `json:"id"`

	// The description of the project category.
	Name string `json:"name"`

	// The URL of the project category.
	Self string `json:"self"`
}
