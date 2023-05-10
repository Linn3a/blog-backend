package models

import "time"

type Passage struct {
	Id 			int
	Title 		string
	Content 	string
	desc		string
	time		time.Time
	CateId		int
}

type Passages []*Passage	