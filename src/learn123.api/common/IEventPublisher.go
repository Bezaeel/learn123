package common

type IEventPublisher interface {
	Publish(event interface{}) error
}
