package types

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
	"strings"
)

type Ticket struct {
	Id         string   `json:"_id"`
	CreatedAt  string   `json:"created_at"`
	Type       string   `json:"type"`
	Subject    string   `json:"subject"`
	AssigneeId float64  `json:"assignee_id"`
	Tags       []string `json:"tags"`
}

type TicketRecords []Ticket

var TicketFields []string = []string{"Id", "CreatedAt", "Type", "Subject", "AssigneeId", "Tags"}

func (t TicketRecords) FindOne(field string, value string) Record {
	var records TicketRecords
	for _, ticket := range t {
		e := reflect.ValueOf(&ticket).Elem()
		for i := 0; i < e.NumField(); i++ {
			varName := e.Type().Field(i).Name
			varValue := e.Field(i).Interface()
			if varName == field {
				stringValue := strings.TrimSuffix(fmt.Sprintln(varValue), "\n")
				if stringValue == value {
					records = append(records, ticket)
				}
			}
		}
	}
	return records
}

func (t TicketRecords) PrintRecord() {
	fmt.Printf("## Search results:\n%s", t.printBasicInfo())
}

func (t TicketRecords) printBasicInfo() string {
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

	for _, record := range t {
		err = tmpl.Execute(&buf, record)
		if err != nil {
			panic(err)
		}
	}

	return buf.String()
}
