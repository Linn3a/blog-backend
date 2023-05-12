package models

type Cate struct {
	Id    uint
	Name  string
	Cover string
	//has many
	Passages Passages
	Tags     Tags
}

type Cates []*Cate
