package zapcore

import (
	"go.uber.org/zap/internal/pool"
)

func encodeError(key string, err error, enc ObjectEncoder) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

type errorGroup interface {
	Errors() []error
}

type errArray []error

func (errs errArray) MarshalLogArray(arr ArrayEncoder) error { _ = "STUB: not implemented"; return nil }

var _errArrayElemPool = pool.New(func() *errArrayElem {
	return &errArrayElem{}
})

type errArrayElem struct{ err error }

func newErrArrayElem(err error) *errArrayElem { _ = "STUB: not implemented"; return nil }

func (e *errArrayElem) MarshalLogArray(arr ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *errArrayElem) MarshalLogObject(enc ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *errArrayElem) Free() { _ = "STUB: not implemented"; return }
