package models

import "time"

type User struct {
	Id 			int
	Username 	string
	Password 	string
	avatar 		string
	desc		string
	gender		string
	birthday	time.Time
	LastLogin 	time.Time
	CreatedTime time.Time
}

type Users []*User