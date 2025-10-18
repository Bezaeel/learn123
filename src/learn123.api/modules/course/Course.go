package course

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"learn123.api/common"
)

// wire module dependencies
func RegisterModule(app *fiber.App, db *gorm.DB, publisher common.IEventPublisher) {
	courseService := NewCourseService(db, publisher)
	NewCourseController(app, courseService)
}
