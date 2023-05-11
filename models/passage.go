package models

import "time"

type Passage struct {
	Id        uint      `json:"id" binding:"-"`
	Title     string    `json:"title"`
	Content   string    `json:"content" binding:"-"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"created_at" binding:"-"`
	CateId    int       `json:"cate_id"`
}

type Passages []*Passage
