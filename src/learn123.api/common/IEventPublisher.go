package common

type IEventPublisher interface {
	Publish(eventType string, event interface{}) error
}
