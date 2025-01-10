package course_test

import (
	"testing"

	"tests"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"learn123.api/modules/course"
)

type CourseServiceShould struct {
	suite.Suite
	sut           course.CourseService
	testDbContext *gorm.DB
}

func TestCourseServiceShouldTestSuite(t *testing.T) {
	suite.Run(t, &CourseServiceShould{})
}

func (uts *CourseServiceShould) SetupTest() {
	uts.testDbContext = tests.SetUp()
	uts.sut = *course.NewCourseService(uts.testDbContext)
}

func (uts *CourseServiceShould) Test_CreateCourse_ShouldCreateCourse_WhenCommandIsValid() {
	// arrange
	newUUID, _ := uuid.NewUUID()
	command := &course.CreateCourseCommand{
		Id:   newUUID,
		Name: "test",
	}

	// act
	expected := uts.sut.CreateCourse(command)

	// assert
	var actual course.CourseEntity
	uts.testDbContext.First(&actual, "Id = ?", command.Id)
	uts.Require().Equal(expected.Value.Id, actual.Id)
}

func (uts *CourseServiceShould) Test_UpdateCourse_ShouldChangeName_WhenCommandIsValid() {
	// arrange
	existingCourseId, _ := uuid.NewUUID()
	uts.testDbContext.Create(&course.CourseEntity{Id: existingCourseId, Name: "first"})
	command := &course.UpdateCourseCommand{
		Name: "test",
	}

	// act
	expected := uts.sut.UpdateCourse(existingCourseId, command)

	// assert
	var actual course.CourseEntity
	uts.testDbContext.First(&actual, "Id = ?", existingCourseId)
	uts.Require().Equal(expected.Value.Name, actual.Name)
}
