package models

import "time"

type User struct {
	UserId     int       `json:"id"`
	Login      string    `json:"login"`
	FirstName  string    `json:"firstName"`
	SecondName string    `json:"secondName"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	RoleId     int       `json:"roleId"`
}
