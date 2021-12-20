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
	fmt.Printf("%+v", ticket)

	templateBody :=
		`       _id:   {{._id}}
	   created_at:   {{.created_at}}
	         type:   {{.type}}
	      subject:   {{.subject}}
		assignee_id:   {{.assigneeId}}
																	 	 										
		`

	tmpl, err := template.New("test").Parse(templateBody)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(&buf, ticket.Attrs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", ticket)
	fmt.Printf("## Search results:\n%s", buf.String())
}

var TicketFields []string = []string{"_id", "created_at", "type", "subject", "assignee_id", "tags"}
