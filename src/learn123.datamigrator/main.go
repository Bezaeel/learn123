package main

import (
	"fmt"

	"learn123.api/modules/course"
	infra "learn123.infrastructure"
	"learn123.infrastructure/database"
)

func main() {
	infra.AddInfrastucture()
	fmt.Println("Running migration...")
	database.DbContext.AutoMigrate(&course.CourseEntity{})
}
