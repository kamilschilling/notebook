package notebook

import "gorm.io/gorm"

type note struct {
	gorm.Model
	Title   string `gorm:"type:varchar(50)"`
	Content string
}

type tag struct {
	gorm.Model
	Name string `gorm:"type:varchar(15)"`
}

type noteTag struct {
	gorm.Model
	NoteID uint
	Note   note
	TagID  uint
	Tag    tag
}
