package types

import (
	"bytes"
	"fmt"
	"html/template"
)

type Ticket struct {
	Attrs map[string]interface{} `json:"_id"`
}

func (t Ticket) Attributes() map[string]interface{} {
	return t.Attrs
}

func (ticket Ticket) Show() {
	var buf bytes.Buffer

	templateBody :=
		`       _id:   {{.Id}}
	   created_at:   {{.CreatedAt}}
	         type:   {{.Type}}
	      subject:   {{.Subject}}
		assignee_id:   {{.AssigneeId}}
																	 	 										
		`

	tmpl, err := template.New("test").Parse(templateBody)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(&buf, ticket)
	if err != nil {
		panic(err)
	}

	fmt.Printf("## Search results:\n%s", buf.String())
}

var TicketFields []string = []string{"Id", "CreatedAt", "Type", "Subject", "AssigneeId", "Tags"}
