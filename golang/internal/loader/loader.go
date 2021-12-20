package loader

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"main.go/internal/types"
)

func LoadData(ctx context.Context, query types.Query) []types.Record {
	var jsonFile *os.File
	var err error
	var records []types.Record
	userPath := "data/users.json"
	ticketPath := "data/tickets.json"
	if os.Getenv("ENVIRONMENT") == "test" {
		userPath = "../../data/users.json"
		ticketPath = "../../data/tickets.json"
	}

	switch query.Dataset {
	case "users":
		jsonFile, err = os.Open(userPath)
		if err != nil {
			panic(err)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		records = loadUsers(byteValue)
	case "tickets":
		jsonFile, err = os.Open(ticketPath)
		if err != nil {
			panic(err)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		records = loadTickets(byteValue)
	default:
		panic(err)
	}

	return records
}

func loadTickets(data []byte) []types.Record {
	var tickets []types.Record
	var ticketsAttributes []map[string]interface{}

	err := json.Unmarshal(data, &ticketsAttributes)
	if err != nil {
		panic(err)
	}

	for _, ticketAttributes := range ticketsAttributes {
		tickets = append(tickets, types.Ticket{Attrs: ticketAttributes})
	}

	return tickets
}

func loadUsers(data []byte) []types.Record {
	var users []types.Record
	var usersAttributes []map[string]interface{}

	err := json.Unmarshal(data, &usersAttributes)
	if err != nil {
		panic(err)
	}

	for _, userAttributes := range usersAttributes {
		users = append(users, types.User{Attrs: userAttributes})
	}

	return users
}
