package app

import (
	"github/four-servings/meonzi/pipe"
)

type EventPublisher interface {
	pipe.PubSub
}

func NewEventPublisher() EventPublisher {
	return pipe.NewPubSub()
}
