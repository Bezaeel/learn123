package course

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseEntity struct {
	gorm.Model
	Id        uuid.UUID `gorm:"column:id;primaryKey`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
