package models

type Tag struct {
	Id       uint
	Name     string
	Color    string
	CateId   uint
	Passages Passages `gorm:"many2many:tag_passages;"`
}

type Tags []*Tag
