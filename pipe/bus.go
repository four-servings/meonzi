package pipe

import (
	"context"
	"errors"
	"reflect"
	"sync"
	"time"
)

var (
	contextType = reflect.TypeOf(new(context.Context)).Elem()
	errorType   = reflect.TypeOf(new(error)).Elem()
)

type ExecuteError interface {
	Error() error
}

type executeErrorImpl struct {
	err error
}

func (e *executeErrorImpl) Error() error {
	return e.err
}

func newExecuteError(err error) ExecuteError {
	return &executeErrorImpl{err}
}

type ExecuteResult interface {
	ExecuteError
	Result(dst interface{}) ExecuteError
}

type executeResultImpl struct {
	ExecuteError
	value interface{}
}

func (e *executeResultImpl) Result(dst interface{}) ExecuteError {
	destination := reflect.ValueOf(dst)
	source := reflect.ValueOf(e.value)

	if reflect.New(source.Type()).Type() != destination.Type() {
		return newExecuteError(errors.New("destination type not match to source type"))
	}

	destination.Elem().Set(source)
	return e.ExecuteError
}

func newExecuteResult(value interface{}, err error) ExecuteResult {
	return &executeResultImpl{
		ExecuteError: newExecuteError(err),
		value:        value,
	}
}

type BusHandlerFunc func(ctx context.Context, value interface{}) (interface{}, error)

type BusHandler interface {
	Wrap() BusHandlerFunc
}

type Bus interface {
	RegistryStrictHandler(targetType interface{}, handler BusHandler) error
	RegistryHandler(targetType interface{}, handler interface{}) error
	Execute(value interface{}) ExecuteResult
	ExecuteContext(ctx context.Context, value interface{}) ExecuteResult
}

type busImpl struct {
	lock    sync.RWMutex
	gateway map[reflect.Type]BusHandlerFunc
	timeout *time.Duration
}

func NewBus() Bus {
	return &busImpl{
		gateway: make(map[reflect.Type]BusHandlerFunc),
	}
}

func NewBusWithTimeout(timeout time.Duration) Bus {
	return &busImpl{
		gateway: make(map[reflect.Type]BusHandlerFunc),
		timeout: &timeout,
	}
}

func (b *busImpl) RegistryHandler(targetType interface{}, handler interface{}) error {
	typ := reflect.TypeOf(targetType)
	if b.exists(typ) {
		return errors.New("already registry value")
	}

	action := reflect.ValueOf(handler)
	actionType := action.Type()
	inCount := actionType.NumIn()
	outCount := actionType.NumOut()

	if action.Kind() != reflect.Func {
		//TODO error message
		return errors.New("")
	}

	if inCount > 2 {
		return errors.New("handler call arguments count overflow")
	} else if inCount == 2 && (contextType != actionType.In(0) || typ != actionType.In(1)) {
		//TODO error message
		return errors.New("")
	} else if inCount == 1 && typ != actionType.In(0) {
		//TODO error message
		return errors.New("")
	}

	if outCount > 2 {
		return errors.New("handler call returns count overflow")
	} else if outCount == 2 && !actionType.Out(1).Implements(errorType) {
		//TODO error message
		return errors.New("")
	} else if outCount == 1 && !actionType.Out(0).Implements(errorType) {
		return errors.New("handler call return must be error")
	}

	b.set(typ, func(ctx context.Context, value interface{}) (ret interface{}, err error) {
		in := make([]reflect.Value, inCount)

		switch inCount {
		case 2:
			in[0] = reflect.ValueOf(ctx)
			in[1] = reflect.ValueOf(value)
		case 1:
			in[0] = reflect.ValueOf(value)
		}

		out := action.Call(in)
		switch outCount {
		case 2:
			ret = out[0].Interface()
			err, _ = out[1].Interface().(error)
		case 1:
			err, _ = out[0].Interface().(error)
		}
		return
	})
	return nil
}

func (b *busImpl) RegistryStrictHandler(targetType interface{}, handler BusHandler) error {
	typ := reflect.TypeOf(targetType)
	if b.exists(typ) {
		return errors.New("already registry value")
	}

	b.set(typ, handler.Wrap())
	return nil
}

func (b *busImpl) Execute(value interface{}) ExecuteResult {
	return b.ExecuteContext(context.Background(), value)
}

func (b *busImpl) ExecuteContext(ctx context.Context, value interface{}) ExecuteResult {
	c := ctx
	if b.timeout != nil {
		var cancel context.CancelFunc
		c, cancel = context.WithTimeout(ctx, *b.timeout)
		defer cancel()
	}

	typ := reflect.TypeOf(value)
	handle, ok := b.get(typ)
	if !ok {
		return newExecuteResult(nil, errors.New("handler not exists"))
	}
	return newExecuteResult(handle(c, value))
}

// critical section method
func (b *busImpl) set(typ reflect.Type, handler BusHandlerFunc) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.gateway[typ] = handler
}

func (b *busImpl) get(typ reflect.Type) (handler BusHandlerFunc, ok bool) {
	b.lock.RLock()
	defer b.lock.RUnlock()
	handler, ok = b.gateway[typ]
	return
}

func (b *busImpl) exists(typ reflect.Type) (ok bool) {
	b.lock.RLock()
	defer b.lock.RUnlock()
	_, ok = b.gateway[typ]
	return
}
