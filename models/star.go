package models

import "time"

type Star struct {
	UserId		uint
	PassageId	int
	CreatedAt	time.Time
}

type Stars []*Star