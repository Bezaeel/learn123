package handlers

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/sirupsen/logrus"
	event "learn123.events"
)

type OrderCreatedHandler struct {
	logger *logrus.Logger
}

func NewOrderCreatedHandler(logger *logrus.Logger) *OrderCreatedHandler {
	return &OrderCreatedHandler{
		logger: logger,
	}
}

func (h *OrderCreatedHandler) Handle(msg *message.Message) ([]*message.Message, error) {
	var orderCreated event.OrderCreated
	if err := json.Unmarshal(msg.Payload, &orderCreated); err != nil {
		return nil, err
	}

	h.logger.WithFields(logrus.Fields{
		"order_id":   orderCreated.Id,
		"order_name": orderCreated.Name,
		"created_by": orderCreated.CreatedBy,
	}).Info("Processing order created event")

	// Add your business logic here

	return nil, nil
}
