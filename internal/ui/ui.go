package ui

import (
	"fmt"

	"main.go/internal/types"
	"main.go/internal/validation"
)

func PromptUser() (types.Query, error) {

	var dataset string
	var field string
	var value string

	dataset = promptForDataset()

	if dataset == "users" {
		field = promptForUserField()
	} else {
		field = promptForTicketField()
	}

	value = promptForSearchValue()

	return types.Query{Dataset: dataset, Field: field, Value: value}, nil
}

func promptForDataset() string {
	for {
		var dataset string
		fmt.Println("Are you searching for tickets or users?")
		fmt.Scanln(&dataset)
		if validation.ValidateDataset(dataset) {
			return dataset
		}
	}
}

func promptForUserField() string {
	for {
		var field string
		fmt.Println("Type user property that you're searching")
		fmt.Scanln(&field)
		if validation.ValidateUserField(field) {
			return field
		}
	}
}

func promptForTicketField() string {
	for {
		var field string
		fmt.Println("Type ticket property that you're searching")
		fmt.Scanln(&field)
		if validation.ValidateTicketField(field) {
			return field
		}
	}
}

func promptForSearchValue() string {
	for {
		var value string
		fmt.Println("Type value you're seacring for")
		fmt.Scanln(&value)
		if validation.ValidateSearchValue(value) {
			return value
		}
	}
}
