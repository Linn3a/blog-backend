package models

import "time"

type Comment struct {
	Id 			int
	Content 	string
	UserId		int
	PassageId	int
	time		time.Time
}

type Comments []*Comment