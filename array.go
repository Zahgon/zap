package zap

import (
	"fmt"
	"time"

	"go.uber.org/zap/zapcore"
)

func Array(key string, val zapcore.ArrayMarshaler) Field {
	_ = "STUB: not implemented"
	return *new(Field)
}

func Bools(key string, bs []bool) Field { _ = "STUB: not implemented"; return *new(Field) }

func ByteStrings(key string, bss [][]byte) Field { _ = "STUB: not implemented"; return *new(Field) }

func Complex128s(key string, nums []complex128) Field {
	_ = "STUB: not implemented"
	return *new(Field)
}

func Complex64s(key string, nums []complex64) Field { _ = "STUB: not implemented"; return *new(Field) }

func Durations(key string, ds []time.Duration) Field { _ = "STUB: not implemented"; return *new(Field) }

func Float64s(key string, nums []float64) Field { _ = "STUB: not implemented"; return *new(Field) }

func Float32s(key string, nums []float32) Field { _ = "STUB: not implemented"; return *new(Field) }

func Ints(key string, nums []int) Field { _ = "STUB: not implemented"; return *new(Field) }

func Int64s(key string, nums []int64) Field { _ = "STUB: not implemented"; return *new(Field) }

func Int32s(key string, nums []int32) Field { _ = "STUB: not implemented"; return *new(Field) }

func Int16s(key string, nums []int16) Field { _ = "STUB: not implemented"; return *new(Field) }

func Int8s(key string, nums []int8) Field { _ = "STUB: not implemented"; return *new(Field) }

func Objects[T zapcore.ObjectMarshaler](key string, values []T) Field {
	_ = "STUB: not implemented"
	return *new(Field)
}

type objects[T zapcore.ObjectMarshaler] []T

func (os objects[T]) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type ObjectMarshalerPtr[T any] interface {
	*T
	zapcore.ObjectMarshaler
}

func ObjectValues[T any, P ObjectMarshalerPtr[T]](key string, values []T) Field {
	_ = "STUB: not implemented"
	return *new(Field)
}

type objectValues[T any, P ObjectMarshalerPtr[T]] []T

func (os objectValues[T, P]) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func Strings(key string, ss []string) Field { _ = "STUB: not implemented"; return *new(Field) }

func Stringers[T fmt.Stringer](key string, values []T) Field {
	_ = "STUB: not implemented"
	return *new(Field)
}

type stringers[T fmt.Stringer] []T

func (os stringers[T]) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func Times(key string, ts []time.Time) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uints(key string, nums []uint) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uint64s(key string, nums []uint64) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uint32s(key string, nums []uint32) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uint16s(key string, nums []uint16) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uint8s(key string, nums []uint8) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uintptrs(key string, us []uintptr) Field { _ = "STUB: not implemented"; return *new(Field) }

func Errors(key string, errs []error) Field { _ = "STUB: not implemented"; return *new(Field) }

type bools []bool

func (bs bools) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type byteStringsArray [][]byte

func (bss byteStringsArray) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type complex128s []complex128

func (nums complex128s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type complex64s []complex64

func (nums complex64s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type durations []time.Duration

func (ds durations) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type float64s []float64

func (nums float64s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type float32s []float32

func (nums float32s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type ints []int

func (nums ints) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type int64s []int64

func (nums int64s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type int32s []int32

func (nums int32s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type int16s []int16

func (nums int16s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type int8s []int8

func (nums int8s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type stringArray []string

func (ss stringArray) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type times []time.Time

func (ts times) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type uints []uint

func (nums uints) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type uint64s []uint64

func (nums uint64s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type uint32s []uint32

func (nums uint32s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type uint16s []uint16

func (nums uint16s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type uint8s []uint8

func (nums uint8s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type uintptrs []uintptr

func (nums uintptrs) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}
