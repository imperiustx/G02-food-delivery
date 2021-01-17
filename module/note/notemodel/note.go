package notemodel

import "fooddlv/common"

const EntityName = "Note"

type Note struct {
	common.SQLModel `json:",inline"`
	Cover           *common.Image  `json:"cover" gorm:"column:cover;"`
	Photos          *common.Images `json:"photos" gorm:"column:photos;"`
	Title           string         `json:"title" gorm:"column:title;"`
	Content         string         `json:"content" gorm:"column:content;"`
}

func (Note) TableName() string {
	return "notes"
}
