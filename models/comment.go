package models

import (
	"time"
)

type Comment struct {
	Id 			uint
	Content 	string
	UserId		uint
	PassageId	uint
	CreatedAt	time.Time
}

type Comments []*Comment

