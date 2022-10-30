package registry

import (
	"bytes"
	"text/template"
)

type Template struct {
	RepoURL string
}

const AUTO_GENERATED_MSG string = `### GRANTED-REGISTRY-SECTION: "{{.RepoURL}}". DO NOT EDIT.
# This section is automatically generated by Granted (https://granted.dev). Manual edits to this section will be overwritten.
# To edit, clone "{{.RepoURL}}", edit granted.yml, and push your changes. You may need to make a pull request depending on the repository settings.
# To stop syncing and remove this section, run 'granted registry remove {{.RepoURL}}
`

func GetAutogeneratedTemplate(repoURL string) string {
	td := Template{RepoURL: repoURL}

	t, err := template.New("autogenerate-msg").Parse(AUTO_GENERATED_MSG)
	if err != nil {
		return ""
	}

	var buf bytes.Buffer
	_ = t.Execute(&buf, td)

	return buf.String()
}
