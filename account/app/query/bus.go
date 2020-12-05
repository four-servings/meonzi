package query

import (
	"errors"
)

type (
	// Bus query bus interface
	Bus interface {
		Handle(query interface{}) interface{}
	}

	busImplement struct {
		findByIDHandler findByIDHandler
		findHandler     findHandler
	}
)

// NewBus create bus instance
func NewBus(query AccountQuery) Bus {
	findByIDHandler := newFindByIDHandler(query)
	findHandler := newFindHandler(query)
	return &busImplement{findByIDHandler, findHandler}
}

// Handle handle given query
func (b *busImplement) Handle(givenQuery interface{}) interface{} {
	switch givenQuery := givenQuery.(type) {
	case *FindByID:
		return b.findByIDHandler.handle(givenQuery)
	case *Find:
		return b.findHandler.handle(givenQuery)
	default:
		panic(errors.New("invalid query"))
	}
}
