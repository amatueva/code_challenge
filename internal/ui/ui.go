package ui

import (
	"fmt"

	"main.go/internal/types"
)

func PromptUser() (types.Query, error) {
	fmt.Println("Type which data type you'd like to use", "tickets", "organizations", "users")

	var dataset string
	fmt.Scanln(&dataset)

	var field string
	fmt.Scanln(&field)

	var value string
	fmt.Scanln(&value)

	return types.Query{Dataset: dataset, Field: field, Value: value}, nil
}
