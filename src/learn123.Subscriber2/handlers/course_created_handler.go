package handlers

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/sirupsen/logrus"
	event "learn123.events"
)

type CourseCreatedHandler struct {
	logger *logrus.Logger
}

func NewCourseCreatedHandler(logger *logrus.Logger) *CourseCreatedHandler {
	return &CourseCreatedHandler{
		logger: logger,
	}
}

func (h *CourseCreatedHandler) Handle(msg *message.Message) ([]*message.Message, error) {
	var courseCreated event.CourseCreated
	if err := json.Unmarshal(msg.Payload, &courseCreated); err != nil {
		return nil, err
	}

	h.logger.WithFields(logrus.Fields{
		"course_id":   courseCreated.Id,
		"course_name": courseCreated.Name,
		"created_by":  courseCreated.CreatedBy,
	}).Info("Processing course created event")

	// Add your business logic here

	return nil, nil
}
