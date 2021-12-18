package types

import (
	"fmt"
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

func (u User) FindOne(field string, value string) {
	e := reflect.ValueOf(&u).Elem()

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varValue := e.Field(i).Interface()
		if varName == field {
			stringValue := strings.TrimSuffix(fmt.Sprintln(varValue), "\n")
			if stringValue == value {
				fmt.Println(u)
			}
		}
	}
}
