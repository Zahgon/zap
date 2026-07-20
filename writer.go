package zap

import (
	"go.uber.org/zap/zapcore"
)

func Open(paths ...string) (zapcore.WriteSyncer, func(), error) {
	_ = "STUB: not implemented"
	return *new(zapcore.WriteSyncer), nil, nil
}

func open(paths []string) ([]zapcore.WriteSyncer, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CombineWriteSyncers(writers ...zapcore.WriteSyncer) zapcore.WriteSyncer {
	_ = "STUB: not implemented"
	return *new(zapcore.WriteSyncer)
}
