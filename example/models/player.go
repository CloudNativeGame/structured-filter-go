package models

type Player struct {
	User  User `json:"user"`
	Level int  `json:"level"`
}

type User struct {
	Name   string `json:"name"`
	IsMale bool   `json:"isMale"`
}
