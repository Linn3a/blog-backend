package models

import "time"

type Star struct {
	UserId		uint
	PassageId	uint
	CreatedAt	time.Time
}

type Stars []*Star