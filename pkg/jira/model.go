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

// Body for POST https://{host}/rest/api/2/issue/{issueIdOrKey}/transitions
//
//go:generate go run ../../main/generator -struct PostTransitionBody -unmarshal
type PostTransitionBody struct {
	// List of issue screen fields to update, specifying the sub-field to
	// update and its value for each field. This field provides a
	// straightforward option when setting a sub-field.
	// When multiple sub-fields or other operations are required, use update.
	// Fields included in here cannot be included in update.
	Fields interface{} `json:"fields,omitempty"`

	// Additional issue history details.
	HistoryMetadata *HistoryMetadata `json:"historyMetadata,omitempty"`

	// Details of issue properties to be add or update.
	Properties []EntityProperty `json:"properties,omitempty"`

	// Details of a transition. Required when performing a transition,
	// optional when creating or editing an issue.
	Transition *IssueTransition `json:"transition,omitempty"`

	// A Map containing the field field name and a list of operations to
	// perform on the issue screen field. Note that fields included in
	// here cannot be included in fields.
	Update interface{} `json:"update,omitempty"`

	// Extra properties of any type may be provided to this object.
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}

//go:generate go run ../../main/generator -struct HistoryMetadata -unmarshal
type HistoryMetadata struct {
	// The activity described in the history record.
	ActivityDescription string `json:"activityDescription,omitempty"`

	// The key of the activity described in the history record.
	ActivityDescriptionKey string `json:"activityDescriptionKey,omitempty"`

	// Details of the user whose action created the history record.
	Actor *HistoryMetadataParticipant `json:"actor,omitempty"`

	// Details of the cause that triggered the creation the history record.
	Cause *HistoryMetadataParticipant `json:"cause,omitempty"`

	// The description of the history record.
	Description string `json:"description,omitempty"`

	// The description key of the history record.
	DescriptionKey string `json:"descriptionKey,omitempty"`

	// The description of the email address associated the history record.
	EmailDescription string `json:"emailDescription,omitempty"`

	// The description key of the email address associated the history record.
	EmailDescriptionKey string `json:"emailDescriptionKey,omitempty"`

	// Additional arbitrary information about the history record.
	ExtraData interface{} `json:"extraData,omitempty"`

	// Details of the system that generated the history record.
	Generator *HistoryMetadataParticipant `json:"generator,omitempty"`

	// The type of the history record.
	Type string `json:"type,omitempty"`

	// Extra properties of any type may be provided to this object.
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}

type EntityProperty struct {
	// The key of the property. Required on create and update.
	Key string `json:"key,omitempty"`

	// The value of the property. Required on create and update.
	Value any `json:"value,omitempty"`
}

//go:generate go run ../../main/generator -struct HistoryMetadataParticipant -unmarshal
type HistoryMetadataParticipant struct {
	// The URL to an avatar for the user or system associated with a history record.
	AvatarUrl string `json:"avatarUrl,omitempty"`

	// The display name of the user or system associated with a history record.
	DisplayName string `json:"displayName,omitempty"`

	// The key of the display name of the user or system associated with a history record.
	DisplayNameKey string `json:"displayNameKey,omitempty"`

	// The ID of the user or system associated with a history record.
	Id string `json:"id,omitempty"`

	// The type of the user or system associated with a history record.
	Type string `json:"type,omitempty"`

	// The URL of the user or system associated with a history record.
	Url string `json:"url,omitempty"`

	// Extra properties of any type may be provided to this object.
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}

// Response from GET https://{host}/rest/api/3/issue/{issueIdOrKey}/transitions
type GetTransitionResponse struct {
	// Expand options that include additional transitions details in the response.
	Expand string `json:"expand,omitempty"`

	// List of issue transitions.
	Transitions []IssueTransition `json:"transitions,omitempty"`
}

//go:generate go run ../../main/generator -struct IssueTransition -unmarshal
type IssueTransition struct {
	// Expand options that include additional transition details in the response.
	Expand string `json:"expand,omitempty"`

	// Details of the fields associated with the issue transition screen.
	// Use this information to populate fields and update in a transition request.
	Fields interface{} `json:"fields,omitempty"`

	// Whether there is a screen associated with the issue transition.
	HasScreen bool `json:"hasScreen,omitempty"`

	// The ID of the issue transition.
	// Required when specifying a transition to undertake.
	Id string `json:"id,omitempty"`

	// Whether the transition is available to be performed.
	IsAvailable bool `json:"isAvailable,omitempty"`

	// Whether the issue has to meet criteria before the issue transition is applied.
	IsConditional bool `json:"isConditional,omitempty"`

	// Whether the issue transition is global, that is,
	// the transition is applied to issues regardless of their status.
	IsGlobal bool `json:"isGlobal,omitempty"`

	// Whether this is the initial issue transition for the workflow.
	IsInitial bool `json:"isInitial,omitempty"`

	Looped bool `json:"looped,omitempty"`

	// The name of the issue transition
	Name string `json:"name,omitempty"`

	// Details of the issue status after the transition.
	To *StatusDetails `json:"to,omitempty"`

	// Extra properties of any type may be provided to this object.
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}

//go:generate go run ../../main/generator -struct StatusDetails  -unmarshal
type StatusDetails struct {
	// The description of the status.
	Description string `json:"description,omitempty"`

	// The URL of the icon used to represent the status.
	IconUrl string `json:"iconUrl,omitempty"`

	// The ID of the status.
	Id string `json:"id,omitempty"`

	// The name of the status.
	Name string `json:"name,omitempty"`

	// The scope of the status.
	Scope *Scope `json:"scope,omitempty"`

	// The URL of the status.
	Self string `json:"self,omitempty"`

	// The category assigned to the status.
	StatusCategory *StatusCategory `json:"statusCategory,omitempty"`

	// Extra properties of any type may be provided to this object.
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}

//go:generate go run ../../main/generator -struct StatusCategory -unmarshal
type StatusCategory struct {
	// The name of the color used to represent the status category.
	ColorName string `json:"colorName,omitempty"`

	// The ID of the status category.
	// Format: int64
	Id int64 `json:"id,omitempty"`

	// The key of the status category.
	Key string `json:"key,omitempty"`

	// The name of the status category.
	Name string `json:"name,omitempty"`

	// The URL of the status category.
	Self string `json:"self,omitempty"`

	// Extra properties of any type may be provided to this object.
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}

//go:generate go run ../../main/generator -struct Scope -unmarshal
type Scope struct {
	// The project the item has scope in.
	Project *ProjectDetails `json:"project,omitempty"`

	// The type of scope.
	// Valid values: PROJECT, TEMPLATE
	Type string `json:"type,omitempty"`

	// Extra properties of any type may be provided to this object.
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}

type ProjectDetails struct {
	// The URLs of the project's avatars.
	AvatarUrls *AvatarUrlsBean `json:"avatarUrls,omitempty"`

	// The ID of the project.
	Id string `json:"id,omitempty"`

	// The key of the project.
	Key string `json:"key,omitempty"`

	// The name of the project.
	Name string `json:"name,omitempty"`

	// The category the project belongs to.
	ProjectCategory *UpdatedProjectCategory `json:"projectCategory,omitempty"`

	// The project type of the project.
	// Valid values: software, service_desk, business
	ProjectTypeKey string `json:"projectTypeKey,omitempty"`

	// The URL of the project details.
	Self string `json:"self,omitempty"`

	// Whether or not the project is simplified.
	Simplified bool `json:"simplified,omitempty"`
}

type AvatarUrlsBean struct {
	// The URL of the item's 16x16 pixel avatar.
	Url16x16 string `json:"16x16,omitempty"`

	// The URL of the item's 24x24 pixel avatar.
	Url24x24 string `json:"24x24,omitempty"`

	// The URL of the item's 32x32 pixel avatar.
	Url32x32 string `json:"32x32,omitempty"`

	// The URL of the item's 48x48 pixel avatar.
	Url48x48 string `json:"48x48,omitempty"`
}

type UpdatedProjectCategory struct {
	// The name of the project category.
	Description string `json:"description,omitempty"`

	// The ID of the project category.
	Id string `json:"id,omitempty"`

	// The description of the project category.
	Name string `json:"name,omitempty"`

	// The URL of the project category.
	Self string `json:"self,omitempty"`
}
