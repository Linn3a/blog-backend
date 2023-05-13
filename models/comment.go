package models

import "time"

type Comment struct {
	Id        uint      `json:"id"`
	Content   string    `json:"content"`
	UserId    uint      `json:"user_id"`
	PassageId uint      `json:"passage_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Comments []*Comment
