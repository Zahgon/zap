package zap

import (
	"go.uber.org/zap/internal/pool"
	"go.uber.org/zap/zapcore"
)

var _errArrayElemPool = pool.New(func() *errArrayElem {
	return &errArrayElem{}
})

func Error(err error) Field { _ = "STUB: not implemented"; return *new(Field) }

func NamedError(key string, err error) Field { _ = "STUB: not implemented"; return *new(Field) }

type errArray []error

func (errs errArray) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type errArrayElem struct {
	error
}

func (e *errArrayElem) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}
