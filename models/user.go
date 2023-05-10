package models

import "time"

type User struct {
	Id 			int
	Username 	string
	Password 	string
	Avatar 		string
	Desc		string
	Gender		string
	IsAdmin		bool
	Birthday	time.Time
	LastLogin 	time.Time
	CreatedTime time.Time
}

type Users []*User