package application

import "github/four-servings/meonzi/local"

type EventPublisher interface {
	local.PubSub
}


func NewEventPublisher() EventPublisher {
	return local.NewPubSub()
}