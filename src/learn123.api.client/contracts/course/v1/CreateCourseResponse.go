package v1

import "github.com/google/uuid"

type CreateCourseResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
