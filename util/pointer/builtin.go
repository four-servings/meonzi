package pointer

import (
	"time"
)

func Time(t time.Time) *time.Time {
	return &t
}

func String(s string) *string {
	return &s
}

func Int(i int) *int {
	return &i
}

func Int8(i int8) *int8 {
	return &i
}

func Int16(i int16) *int16 {
	return &i
}

func Int32(i int32) *int32 {
	return &i
}

func Int64(i int64) *int64 {
	return &i
}

func Uint(ui uint) *uint {
	return &ui
}

func Uint8(ui uint8) *uint8 {
	return &ui
}

func Uint16(ui uint16) *uint16 {
	return &ui
}

func Uint32(ui uint32) *uint32 {
	return &ui
}

func Uint64(ui uint64) *uint64 {
	return &ui
}

func Float32(f float32) *float32 {
	return &f
}

func Float64(f float64) *float64 {
	return &f
}