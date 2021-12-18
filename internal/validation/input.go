package validation

import "main.go/internal/types"

func ValidateDataset(dataset string) bool {
	if dataset == "users" || dataset == "tickets" {
		return true
	}
	return false
}

func ValidateUserField(field string) bool {
	for _, userField := range types.UserFields {
		if field == userField {
			return true
		}
	}
	return false
}

func ValidateTicketField(field string) bool {
	for _, ticketField := range types.TicketFields {
		if field == ticketField {
			return true
		}
	}
	return false
}

func ValidateSearchValue(value string) bool {
	if len(value) > 0 {
		return true
	}
	return false
}
