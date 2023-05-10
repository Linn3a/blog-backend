package models

import "time"

type Comment struct {
	Id 			int
	Comment 	string
	UserId		int
	PassageId	int
	time		time.Time
}

type Comments []*Comment