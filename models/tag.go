package models


type Tag struct {
	Id 			int
	Name 		string
	color		string
	CateId		int
}

type Tags []*Tag