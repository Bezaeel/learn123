package course

import (
	"gorm.io/gorm"
	ext "learn123.api/common/extensions"
)

type ICourseService interface {
	CreateCourse(command *CreateCourseCommand) ext.Result[*CourseEntity]
}

type CourseService struct {
	dbContext *gorm.DB
}

func (service *CourseService) CreateCourse(command *CreateCourseCommand) ext.Result[*CourseEntity] {

	result := service.dbContext.Create(command.ToEntity())
	if result.Error != nil {
		return ext.Result[*CourseEntity]{Err: result.Error}
	}

	// publish CourseCreated event
	return ext.Result[*CourseEntity]{Value: command.ToEntity()}
}

func NewCourseService(dbContext *gorm.DB) *CourseService {
	return &CourseService{
		dbContext: dbContext,
	}
}
