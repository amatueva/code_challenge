package types

// type Record interface {
// 	FindOne(field string, value string) Record
// 	PrintRecord()
// }

type Record interface {
	Attributes() map[string]interface{}
	Show()
}
