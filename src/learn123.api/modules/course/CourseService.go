package course

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	ext "learn123.api/common/extensions"
)

type ICourseService interface {
	CreateCourse(command *CreateCourseCommand) ext.Result[*CourseEntity]
	UpdateCourse(id uuid.UUID, command *UpdateCourseCommand) ext.Result[*CourseEntity]
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

func (service *CourseService) UpdateCourse(id uuid.UUID, command *UpdateCourseCommand) ext.Result[*CourseEntity] {
	var result CourseEntity
	err := service.dbContext.Model(&result).
		Clauses(clause.Returning{}).
		Where("Id = ?", id).
		Updates(CourseEntity{Name: command.Name}).Error

	if err != nil {
		return ext.Result[*CourseEntity]{Err: err}
	}

	return ext.Result[*CourseEntity]{Value: &result}
}

func NewCourseService(dbContext *gorm.DB) *CourseService {
	return &CourseService{
		dbContext: dbContext,
	}
}
