package types

type User struct {
	Id        float64 `json:"_id"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"created_at"`
	Verified  bool    `json:"verified"`
}
