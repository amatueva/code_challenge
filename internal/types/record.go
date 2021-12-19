package types

type Record interface {
	FindOne(field string, value string) Record
	PrintRecord()
}
