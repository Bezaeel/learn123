package tests

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"learn123.api/modules/course"
)

func SetUp() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("unable to setup db")
	}

	db.AutoMigrate(&course.CourseEntity{})
	return db
}
