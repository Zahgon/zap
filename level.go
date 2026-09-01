package zap

import (
	"sync/atomic"

	"go.uber.org/zap/internal"
	"go.uber.org/zap/zapcore"
)

const (
	DebugLevel = zapcore.DebugLevel

	InfoLevel = zapcore.InfoLevel

	WarnLevel = zapcore.WarnLevel

	ErrorLevel = zapcore.ErrorLevel

	DPanicLevel = zapcore.DPanicLevel

	PanicLevel = zapcore.PanicLevel

	FatalLevel = zapcore.FatalLevel
)

type LevelEnablerFunc func(zapcore.Level) bool

func (f LevelEnablerFunc) Enabled(lvl zapcore.Level) bool { _ = "STUB: not implemented"; return false }

type AtomicLevel struct {
	l *atomic.Int32
}

var _ internal.LeveledEnabler = AtomicLevel{}

func NewAtomicLevel() AtomicLevel { _ = "STUB: not implemented"; return *new(AtomicLevel) }

func NewAtomicLevelAt(l zapcore.Level) AtomicLevel {
	_ = "STUB: not implemented"
	return *new(AtomicLevel)
}

func ParseAtomicLevel(text string) (AtomicLevel, error) {
	_ = "STUB: not implemented"
	return *new(AtomicLevel), nil
}

func (lvl AtomicLevel) Enabled(l zapcore.Level) bool { _ = "STUB: not implemented"; return false }

func (lvl AtomicLevel) Level() zapcore.Level { _ = "STUB: not implemented"; return *new(zapcore.Level) }

func (lvl AtomicLevel) SetLevel(l zapcore.Level) { _ = "STUB: not implemented"; return }

func (lvl AtomicLevel) String() string { _ = "STUB: not implemented"; return "" }

func (lvl *AtomicLevel) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (lvl AtomicLevel) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
