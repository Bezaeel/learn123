package course

import "github.com/google/uuid"

type CreateCourseCommand struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UpdateCourseCommand struct {
	Name string `json:"name"`
}