package models

import "time"

type Star struct {
	UserId    uint      `json:"user_id"`
	PassageId uint      `json:"passage_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Stars []*Star
