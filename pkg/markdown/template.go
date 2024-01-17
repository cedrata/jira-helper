package markdown

import (
	"bytes"
	"os"
	"sync"
	"text/template"

	"github.com/cedrata/jira-helper/pkg/jira"
)

const DocTemplate = `
# {{.Summary}} {{.Key}}

{{.Description}}

### Introduction

### Conclusion

### Bibliography
`

const PresTemplate = `
# {{.Summary}} {{.Key}}

{{.Description}}

--------------------

# Introduction

--------------------

# Conclusion

`

const testTableTemplate = `
Key | Summary
----|--------:
{{range .Tests}}
{{- .Key -}} | {{ .Summary }}
{{end}}
`

const summaryTemplate = `# {{.Name}}
{{range .Issues}}
## {{.Key}} - {{.Summary}}

### description
{{.Description}}

### conclusion/tests

{{end}}
`

type Summary struct {
	Issues jira.Issues
	Name   string
}

const DocFileSrc = "doc.md"
const PresFileSrc = "presentation.md"

var tmpl *template.Template
var tmplOnce sync.Once

func TemplatleInstance() *template.Template {
	tmplOnce.Do(func() {
		tmpl = template.New("test")
	})
	return tmpl
}

func WriteStub(story jira.Issue, file string, template string) error {
	parsed, err := TemplatleInstance().Parse(template)
	if err != nil {
		return err
	}

	stream, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0622)
	if err != nil {
		return err
	}

	return parsed.Execute(stream, story)
}

func WriteTestTable(test jira.TestList, file string) error {
	parsed, err := TemplatleInstance().Parse(testTableTemplate)
	if err != nil {
		return err
	}

	stream, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0622)
	if err != nil {
		return err
	}

	return parsed.Execute(stream, test)
}

func GenerateIssuesMarkdown(summary *Summary) (*bytes.Buffer, error) {
	parsed, err := TemplatleInstance().Parse(summaryTemplate)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	parsed.Execute(buf, summary)
	return buf, nil
}
