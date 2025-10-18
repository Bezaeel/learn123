package learn123events

import "github.com/google/uuid"

type OrderCreated struct {
	Id        uuid.UUID
	Name      string
	CreatedBy string
}
