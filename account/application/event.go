package application

import (
	"github.com/google/uuid"
	"github/four-servings/meonzi/ent"
)

type (
	UserRegisteredEvent struct {
		Id uuid.UUID
		Account ent.Account
	}

	UserUpdatedEvent struct {
		Id uuid.UUID
		Account ent.Account
	}

	UserRemovedEvent struct {
		Id uuid.UUID
	}
)