package models


type Tag struct {
	Id 			int
	Name 		string
	Color		string
	CateId		int
}

type Tags []*Tag