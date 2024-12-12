package models

type Player struct {
	Id     int      `json:"id"`
	User   User     `json:"user"`
	Level  int      `json:"level"`
	Ids    []int    `json:"ids"`
	Labels []string `json:"labels"`
}

type User struct {
	Name   string `json:"name"`
	IsMale bool   `json:"isMale"`
}
