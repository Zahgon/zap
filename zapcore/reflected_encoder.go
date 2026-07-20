package zapcore

import (
	"io"
)

type ReflectedEncoder interface {
	Encode(interface{}) error
}

func defaultReflectedEncoder(w io.Writer) ReflectedEncoder {
	_ = "STUB: not implemented"
	return *new(ReflectedEncoder)
}
