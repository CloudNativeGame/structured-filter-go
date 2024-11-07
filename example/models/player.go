package models

type Player struct {
	User User `json:"user"`
}

type User struct {
	Name   string `json:"name"`
	IsMale bool   `json:"isMale"`
}
