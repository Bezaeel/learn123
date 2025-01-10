package course

import (
	"time"

	"github.com/google/uuid"
)

type CourseEntity struct {
	Id             uuid.UUID `gorm:"column:id;primaryKey"`
	Name           string    `gorm:"column:name"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	LastModifiedAt time.Time `gorm:"column:updated_at"`
}

func (c *CourseEntity) TableName() string {
	return "course"
}
