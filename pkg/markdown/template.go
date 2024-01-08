package markdown

import (
	"os"
	"text/template"

	"github.com/cedrata/jira-helper/pkg/jira"
)

const DocTemplate =`
# {{.Summary}} {{.Key}}

{{.Description}}

### Introduction

### Conclusion

### Bibliography
`

const PresTemplate =`
# {{.Summary}} {{.Key}}

{{.Description}}

--------------------

# Introduction

--------------------

# Conclusion

`

const DocFileSrc = "doc.md"
const PresFileSrc = "presentation.md"


var tmpl *template.Template

func TemplatleInstance() *template.Template{
	if tmpl == nil {
		tmpl = template.New("test")
	}
	return tmpl
}

func WriteStub(story jira.Issue, file string, template string) error {
	parsed, err := TemplatleInstance().Parse(template)
	if err != nil {
		return err
	}

	stream, err := os.OpenFile(file, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0622)
	if err != nil {
		return err
	}

	return parsed.Execute(stream, story)
}