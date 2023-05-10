package models

import "time"

type Star struct {
	UserId		int
	PassageId	int
	time		time.Time
}

type Stars []*Star