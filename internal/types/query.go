package types

type Result struct {
	User   User
	Ticket Ticket
}

type Query struct {
	Dataset string
	Field   string
	Value   string
}
