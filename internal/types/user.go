package types

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
	"strings"
)

type User struct {
	Id        float64 `json:"_id"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"created_at"`
	Verified  bool    `json:"verified"`
}

var UserFields []string = []string{"Id", "Name", "CreatedAt", "Verified"}

type UserRecords []User

func (u UserRecords) FindOne(field string, value string) Record {
	var records UserRecords

	for _, user := range u {
		e := reflect.ValueOf(&user).Elem()

		for i := 0; i < e.NumField(); i++ {
			varName := e.Type().Field(i).Name
			varValue := e.Field(i).Interface()
			if varName == field {
				stringValue := strings.TrimSuffix(fmt.Sprintln(varValue), "\n")
				if stringValue == value {
					records = append(records, user)
				}
			}
		}
	}
	return records
}

func (u UserRecords) PrintRecord() {
	fmt.Printf("## Search results:\n%s", u.PrintBasicInfo())
}

func (u UserRecords) PrintBasicInfo() string {
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

	for _, record := range u {
		err = tmpl.Execute(&buf, record)
		if err != nil {
			panic(err)
		}
	}

	return buf.String()
}
