package models

type Tag struct {
	Id 			uint
	Name 		string
	Color		string
	CateId		int
}

type Tags []*Tag