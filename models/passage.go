package models

import "time"

type Passage struct {
	Id        uint      `json:"id" binding:"-"`
	Title     string    `json:"title"`
	Content   string    `json:"content" binding:"-"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"created_at" binding:"-"`
	//has many
	Comments Comments
	//many to many
	Tags   Tags `gorm:"many2many:tag_passages;"`
	CateId uint
}

type Passages []*Passage
