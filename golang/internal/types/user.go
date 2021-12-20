package types

import (
	"bytes"
	"fmt"
	"html/template"
)

type User struct {
	Attrs map[string]interface{} `json:"_id"`
}

func (u User) Attributes() map[string]interface{} {
	return u.Attrs
}

var UserFields []string = []string{"_id", "name", "created_at", "verified"}

func (u User) Show() {
	var buf bytes.Buffer

	templateBody :=
		`                   _id: {{._id}}
	    created_at: {{.created_at}}
	          type: {{.name}}
						
`
	tmpl, err := template.New("test").Parse(templateBody)

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(&buf, u.Attrs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("## Search results:\n%s", buf.String())
}
