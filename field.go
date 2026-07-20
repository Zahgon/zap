package zap

import (
	"fmt"
	"math"
	"time"

	"go.uber.org/zap/zapcore"
)

type Field = zapcore.Field

var (
	_minTimeInt64 = time.Unix(0, math.MinInt64)
	_maxTimeInt64 = time.Unix(0, math.MaxInt64)
)

func Skip() Field { _ = "STUB: not implemented"; return *new(Field) }

func nilField(key string) Field { _ = "STUB: not implemented"; return *new(Field) }

func Binary(key string, val []byte) Field { _ = "STUB: not implemented"; return *new(Field) }

func Bool(key string, val bool) Field { _ = "STUB: not implemented"; return *new(Field) }

func Boolp(key string, val *bool) Field { _ = "STUB: not implemented"; return *new(Field) }

func ByteString(key string, val []byte) Field { _ = "STUB: not implemented"; return *new(Field) }

func Complex128(key string, val complex128) Field { _ = "STUB: not implemented"; return *new(Field) }

func Complex128p(key string, val *complex128) Field { _ = "STUB: not implemented"; return *new(Field) }

func Complex64(key string, val complex64) Field { _ = "STUB: not implemented"; return *new(Field) }

func Complex64p(key string, val *complex64) Field { _ = "STUB: not implemented"; return *new(Field) }

func Float64(key string, val float64) Field { _ = "STUB: not implemented"; return *new(Field) }

func Float64p(key string, val *float64) Field { _ = "STUB: not implemented"; return *new(Field) }

func Float32(key string, val float32) Field { _ = "STUB: not implemented"; return *new(Field) }

func Float32p(key string, val *float32) Field { _ = "STUB: not implemented"; return *new(Field) }

func Int(key string, val int) Field { _ = "STUB: not implemented"; return *new(Field) }

func Intp(key string, val *int) Field { _ = "STUB: not implemented"; return *new(Field) }

func Int64(key string, val int64) Field { _ = "STUB: not implemented"; return *new(Field) }

func Int64p(key string, val *int64) Field { _ = "STUB: not implemented"; return *new(Field) }

func Int32(key string, val int32) Field { _ = "STUB: not implemented"; return *new(Field) }

func Int32p(key string, val *int32) Field { _ = "STUB: not implemented"; return *new(Field) }

func Int16(key string, val int16) Field { _ = "STUB: not implemented"; return *new(Field) }

func Int16p(key string, val *int16) Field { _ = "STUB: not implemented"; return *new(Field) }

func Int8(key string, val int8) Field { _ = "STUB: not implemented"; return *new(Field) }

func Int8p(key string, val *int8) Field { _ = "STUB: not implemented"; return *new(Field) }

func String(key string, val string) Field { _ = "STUB: not implemented"; return *new(Field) }

func Stringp(key string, val *string) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uint(key string, val uint) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uintp(key string, val *uint) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uint64(key string, val uint64) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uint64p(key string, val *uint64) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uint32(key string, val uint32) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uint32p(key string, val *uint32) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uint16(key string, val uint16) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uint16p(key string, val *uint16) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uint8(key string, val uint8) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uint8p(key string, val *uint8) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uintptr(key string, val uintptr) Field { _ = "STUB: not implemented"; return *new(Field) }

func Uintptrp(key string, val *uintptr) Field { _ = "STUB: not implemented"; return *new(Field) }

func Reflect(key string, val interface{}) Field { _ = "STUB: not implemented"; return *new(Field) }

func Namespace(key string) Field { _ = "STUB: not implemented"; return *new(Field) }

func Stringer(key string, val fmt.Stringer) Field { _ = "STUB: not implemented"; return *new(Field) }

func Time(key string, val time.Time) Field { _ = "STUB: not implemented"; return *new(Field) }

func Timep(key string, val *time.Time) Field { _ = "STUB: not implemented"; return *new(Field) }

func Stack(key string) Field { _ = "STUB: not implemented"; return *new(Field) }

func StackSkip(key string, skip int) Field { _ = "STUB: not implemented"; return *new(Field) }

func Duration(key string, val time.Duration) Field { _ = "STUB: not implemented"; return *new(Field) }

func Durationp(key string, val *time.Duration) Field { _ = "STUB: not implemented"; return *new(Field) }

func Object(key string, val zapcore.ObjectMarshaler) Field {
	_ = "STUB: not implemented"
	return *new(Field)
}

func Inline(val zapcore.ObjectMarshaler) Field { _ = "STUB: not implemented"; return *new(Field) }

func Dict(key string, val ...Field) Field { _ = "STUB: not implemented"; return *new(Field) }

func dictField(key string, val []Field) Field { _ = "STUB: not implemented"; return *new(Field) }

type dictObject []Field

func (d dictObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func DictObject(val ...Field) zapcore.ObjectMarshaler {
	_ = "STUB: not implemented"
	return *new(zapcore.ObjectMarshaler)
}

type anyFieldC[T any] func(string, T) Field

func (f anyFieldC[T]) Any(key string, val any) Field { _ = "STUB: not implemented"; return *new(Field) }

func Any(key string, value interface{}) Field { _ = "STUB: not implemented"; return *new(Field) }
