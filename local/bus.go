package local

import (
	"context"
	"errors"
	"reflect"
	"time"
)

type BusHandlerFunc func(ctx context.Context, value interface{}) error

type BusHandler interface {
	Wrap() BusHandlerFunc
}

type Bus interface {
	Registry(value interface{}, handler BusHandler) error
	Execute(value interface{}) error
	ExecuteContext(ctx context.Context, value interface{}) error
}

type busImpl struct {
	gateway map[reflect.Type]BusHandlerFunc
	timeout time.Duration
}

func NewBus(timeout time.Duration) Bus {
	return &busImpl{
		gateway: make(map[reflect.Type]BusHandlerFunc),
		timeout: timeout,
	}
}

func (b *busImpl) Registry(value interface{}, handler BusHandler) error {
	typ := reflect.TypeOf(value)
	_, ok := b.gateway[typ]
	if ok {
		return errors.New("already registry value")
	}
	b.gateway[typ] = handler.Wrap()
	return nil
}

func (b *busImpl) Execute(value interface{}) error {
	c, cancel := context.WithTimeout(context.Background(), b.timeout)
	defer cancel()
	return b.ExecuteContext(c, value)
}

func (b *busImpl) ExecuteContext(ctx context.Context, value interface{}) error {
	typ := reflect.TypeOf(value)
	handle, ok := b.gateway[typ]
	if ok {
		return errors.New("")
	}
	return handle(ctx, value)
}