package course

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// wire module dependencies
func RegisterModule(app *fiber.App, db *gorm.DB) {
	courseService := NewCourseService(db)
	NewCourseController(app, courseService)
}
