package search

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"main.go/internal/types"
)

func Search(ctx context.Context, query types.Query) {
	tickets := loadTickets(ctx)
	users := loadUsers(ctx)
	fmt.Println(users, tickets)
	// do the actual search
}

func loadTickets(ctx context.Context) []types.Ticket {
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

func loadUsers(ctx context.Context) []types.User {
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

// func findOne(query types.Query) Record {
// 	if index[query] != nil {
// 		return index[query][0]
// 	}
// 	return nil
// }
