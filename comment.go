package models

import "time"

type Comment struct {
	Id 			uint
	Content 	string
	UserId		int
	PassageId	int
	CreatedAt	time.Time
}

type Comments []*Comment