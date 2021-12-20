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

var UserFields []string = []string{"Id", "Name", "CreatedAt", "Verified"}

func (u User) Show() {
	var buf bytes.Buffer

	templateBody :=
		`                   _id: {{.Id}}
	    created_at: {{.CreatedAt}}
	          type: {{.Name}}
						
`
	tmpl, err := template.New("test").Parse(templateBody)

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(&buf, u)
	if err != nil {
		panic(err)
	}

	fmt.Printf("## Search results:\n%s", buf.String())
}
