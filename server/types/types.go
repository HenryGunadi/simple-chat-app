package types

type UserStore interface {
}

type AuthStore interface{}

type User struct {
	ID       string `json:"ID"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}