package zapfield

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Str[K ~string, V ~string](k K, v V) zap.Field {
	_ = "STUB: not implemented"
	return *new(zap.Field)
}

type stringArray[T ~string] []T

func (a stringArray[T]) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func Strs[K ~string, V ~[]S, S ~string](k K, v V) zap.Field {
	_ = "STUB: not implemented"
	return *new(zap.Field)
}
