package models

type Tag struct {
	Id       uint     `json:"id"`
	Name     string   `json:"name"`
	Color    string   `json:"color"`
	CateId   uint     `json:"cate_id"`
	Passages Passages `gorm:"many2many:tag_passages;"`
}

type Tags []*Tag
