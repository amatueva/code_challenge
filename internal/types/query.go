package types

type Database struct {
	Users   []User
	Tickets []Ticket
}

type Query struct {
	Dataset string
	Field   string
	Value   interface{}
}
