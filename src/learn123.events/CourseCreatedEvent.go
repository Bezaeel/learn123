package learn123events

import "github.com/google/uuid"

type CourseCreated struct {
	Id        uuid.UUID
	Name      string
	CreatedBy string
}
