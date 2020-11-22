package query

import (
	"errors"
	"github/four-servings/meonzi/accounts/infra"
)

type (
	// Bus query bus interface
	Bus interface {
		Handle(query interface{}) interface{}
	}

	busImplement struct {
		findByIDHandler findByIDHandler
	}
)

// NewBus create bus instance
func NewBus(repository infra.AccountRepository) Bus {
	findByIDHandler := newFindByIDHandler(repository)
	return &busImplement{findByIDHandler}
}

// Handle handle given query
func (b *busImplement) Handle(givenQuery interface{}) interface{} {
	switch givenQuery := givenQuery.(type) {
	case *FindByID:
		return b.findByIDHandler.handle(givenQuery)
	default:
		panic(errors.New("invalid query"))
	}
}
