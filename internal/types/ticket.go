package types

import (
	"fmt"
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

func (t TicketRecords) FindOne(field string, value string) {
	e := reflect.ValueOf(&t).Elem()

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varValue := e.Field(i).Interface()
		if varName == field {
			stringValue := strings.TrimSuffix(fmt.Sprintln(varValue), "\n")
			if stringValue == value {
				fmt.Println(t)
			}
		}
	}
}
