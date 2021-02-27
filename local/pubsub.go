package local

import (
	"errors"
	"reflect"
	"sync"
)

type PubSub interface {
	Publish(value interface{})
	Subscribe(handler interface{}) error
}

//TODO exists check
type subs struct {
	lock    sync.RWMutex
	handler []interface{}
}

func (s *subs) mux(value interface{}) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for i := range s.handler {
		action := reflect.ValueOf(s.handler[i])
		go action.Call([]reflect.Value{reflect.ValueOf(value)})
	}
}

func (s *subs) append(handler interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.handler = append(s.handler, handler)
}

type values struct {
	lock  sync.RWMutex
	value []interface{}
}

func (v *values) append(value ...interface{}) {
	v.lock.Lock()
	defer v.lock.Unlock()
	v.value = append(v.value, value...)
}

func (v *values) get(i int) (res interface{}) {
	v.lock.RLock()
	defer v.lock.RUnlock()
	res = v.value[i]
	return
}

func (v *values) count() (cnt int) {
	v.lock.RLock()
	defer v.lock.RUnlock()
	cnt = len(v.value)
	return
}

// TODO max value count
type pubsubImpl struct {
	valueLock sync.RWMutex
	value     map[reflect.Type]*values

	subLock sync.RWMutex
	sub     map[reflect.Type]*subs
}

func NewPubSub() PubSub {
	return &pubsubImpl{
		value: make(map[reflect.Type]*values),
		sub:   make(map[reflect.Type]*subs),
	}
}

func (ps *pubsubImpl) Publish(value interface{}) {
	go ps.appendValue(value)
	go ps.callSubscribers(value)
}

func (ps *pubsubImpl) Subscribe(handler interface{}) error {
	return ps.appendSub(handler)
}

// critical section
func (ps *pubsubImpl) callSubscribers(value interface{}) {
	typ := reflect.TypeOf(value)

	ps.getSafeSubs(typ).
		mux(value)
}

func (ps *pubsubImpl) appendSub(handler interface{}) error {
	handlerType := reflect.TypeOf(handler)
	if handlerType.Kind() != reflect.Func {
		//TODO error message
		return errors.New("")
	} else if handlerType.NumIn() != 1 {
		//TODO error message
		return errors.New("")
	}

	typ := handlerType.In(0)

	sub := ps.getSafeSubs(typ)
	val := ps.getSafeValues(typ)
	seek := 0
	action := reflect.ValueOf(handler)
	for seek < val.count() {
		action.Call([]reflect.Value{reflect.ValueOf(val.get(seek))})
		seek++
	}
	sub.append(handler)
	return nil
}

func (ps *pubsubImpl) getSafeSubs(typ reflect.Type) (res *subs) {
	if ps.existsSubs(typ) {
		res = ps.getSubs(typ)
	} else {
		res = ps.createSubs(typ)
	}

	return
}

func (ps *pubsubImpl) existsSubs(typ reflect.Type) (ok bool) {
	ps.subLock.RLock()
	defer ps.subLock.RUnlock()
	_, ok = ps.sub[typ]
	return
}

func (ps *pubsubImpl) getSubs(typ reflect.Type) (res *subs) {
	ps.subLock.RLock()
	defer ps.subLock.RUnlock()
	res = ps.sub[typ]
	return
}

func (ps *pubsubImpl) createSubs(typ reflect.Type) (res *subs) {
	res = &subs{}
	ps.subLock.Lock()
	defer ps.subLock.Unlock()
	ps.sub[typ] = res
	return
}

func (ps *pubsubImpl) appendValue(value interface{}) {
	typ := reflect.TypeOf(value)

	ps.getSafeValues(typ).
		append(value)
}

func (ps *pubsubImpl) getSafeValues(typ reflect.Type) (res *values) {
	if ps.existsValues(typ) {
		res = ps.getValues(typ)
	} else {
		res = ps.createValues(typ)
	}

	return
}

func (ps *pubsubImpl) existsValues(typ reflect.Type) (ok bool) {
	ps.valueLock.RLock()
	defer ps.valueLock.RUnlock()
	_, ok = ps.value[typ]
	return
}

func (ps *pubsubImpl) getValues(typ reflect.Type) (res *values) {
	ps.valueLock.RLock()
	defer ps.valueLock.RUnlock()
	res = ps.value[typ]
	return
}

func (ps *pubsubImpl) createValues(typ reflect.Type) (res *values) {
	res = &values{}
	ps.valueLock.Lock()
	defer ps.valueLock.Unlock()
	ps.value[typ] = res
	return
}
