package loader

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"main.go/internal/types"
)

func LoadTickets(ctx context.Context) []types.Ticket {
	jsonFile, err := os.Open("data/tickets.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tickets []types.Ticket

	err = json.Unmarshal(byteValue, &tickets)
	if err != nil {
		panic(err)
	}

	return tickets
}

func LoadUsers(ctx context.Context) []types.User {
	jsonFile, err := os.Open("data/users.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users []types.User

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		panic(err)
	}

	return users
}
