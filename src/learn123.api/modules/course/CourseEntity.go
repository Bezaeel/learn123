package course

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseEntity struct {
	gorm.Model
	Id   uuid.UUID `gorm:"column:id;primaryKey`
	Name string    `gorm:"column:name"`
}

func (c *CourseEntity) TableName() string {
	return "course"
}
