package loader

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"main.go/internal/types"
)

func LoadData(ctx context.Context, query types.Query) types.Record {
	var jsonFile *os.File
	var err error
	var records types.Record

	switch query.Dataset {
	case "users":
		jsonFile, err = os.Open("data/users.json")
		byteValue, _ := ioutil.ReadAll(jsonFile)
		records = loadUsers(byteValue)
	case "tickets":
		jsonFile, err = os.Open("data/tickets.json")
		byteValue, _ := ioutil.ReadAll(jsonFile)
		records = loadTickets(byteValue)
	default:
	}

	if err != nil {
		panic(err)
	}

	return records
}

func loadTickets(data []byte) types.Record {
	var tickets types.TicketRecords

	err := json.Unmarshal(data, &tickets)
	if err != nil {
		panic(err)
	}

	return tickets
}

func loadUsers(data []byte) types.Record {
	var users types.UserRecords

	err := json.Unmarshal(data, &users)
	if err != nil {
		panic(err)
	}

	return users
}
