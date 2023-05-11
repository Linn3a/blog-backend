package models

import "time"

type Passage struct {
	Id 			uint
	Title 		string
	Content 	string
	desc		string
	CreatedAt	time.Time
	CateId		int
}

type Passages []*Passage	