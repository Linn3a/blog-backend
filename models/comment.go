package models

import "time"

type Comment struct {
	Id        uint
	Content   string
	UserId    uint `json:"user_id"`
	PassageId uint `json:"passage_id"`
	CreatedAt time.Time
}

type Comments []*Comment
