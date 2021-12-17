package types

type Ticket struct {
	Id         string   `json:"_id"`
	CreatedAt  string   `json:"created_at"`
	Type       string   `json:"type"`
	Subject    string   `json:"subject"`
	AssigneeId float64  `json:"assignee_id"`
	Tags       []string `json:"tags"`
}
