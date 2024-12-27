package course

import (
	"github.com/google/uuid"
	courseV1 "learn123.api.client/contracts/course/v1"
)

func (command *CreateCourseCommand) ToEntity() *CourseEntity {
	return &CourseEntity{
		Id:   uuid.New(),
		Name: command.Name,
	}
}

func (entity *CourseEntity) ToResponse() *courseV1.CreateCourseResponse {
	return &courseV1.CreateCourseResponse{
		Id:   entity.Id,
		Name: entity.Name,
	}
}
