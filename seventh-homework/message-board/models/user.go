package models

type User struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}
