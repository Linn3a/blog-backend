package models

type Cate struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Cover string `json:"cover"`
	//has many
	Passages Passages `json:"passages"`
	Tags     Tags     `json:"tags"`
}

type Cates []*Cate
