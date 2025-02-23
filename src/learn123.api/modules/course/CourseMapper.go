package course

import (
	courseV1 "learn123.api.client/contracts/course/v1"
	event "learn123.events"
)

func (command *CreateCourseCommand) ToEvent(createdBy string) *event.CourseCreated {
	return &event.CourseCreated{
		Id:        command.Id,
		Name:      command.Name,
		CreatedBy: createdBy,
	}
}

func (command *CreateCourseCommand) ToEntity() *CourseEntity {
	return &CourseEntity{
		Id:   command.Id,
		Name: command.Name,
	}
}

func (entity *CourseEntity) ToResponse() *courseV1.CreateCourseResponse {
	return &courseV1.CreateCourseResponse{
		Id:   entity.Id,
		Name: entity.Name,
	}
}
