package course

import (
	"github.com/gofiber/fiber/v2"
)

type CourseController struct {
	router        fiber.Router
	courseService ICourseService
}

func NewCourseController(app *fiber.App, _courseService ICourseService) {
	controller := &CourseController{
		router:        app.Group("/api/course/v1"),
		courseService: _courseService,
	}

	controller.Ask()
	controller.CreateCourse()
}

func (c *CourseController) Ask() {
	c.router.Get("/api/v1/ask", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON("Ask Talabi..")
	})
}

func (c *CourseController) CreateCourse() {
	c.router.Post("/new", func(ctx *fiber.Ctx) error {
		p := new(CreateCourseCommand)
		_ = ctx.BodyParser(p)

		result := c.courseService.CreateCourse(p)

		if result.IsFailure() {
			return ctx.Status(fiber.StatusInternalServerError).JSON(result.Err)
		}

		return ctx.Status(fiber.StatusOK).JSON(result.Value.ToResponse())
	})
}
