package zapcore

import (
	"errors"
)

var errUnmarshalNilLevel = errors.New("can't unmarshal a nil *Level")

type Level int8

const (
	DebugLevel Level = iota - 1

	InfoLevel

	WarnLevel

	ErrorLevel

	DPanicLevel

	PanicLevel

	FatalLevel

	_minLevel = DebugLevel
	_maxLevel = FatalLevel

	InvalidLevel = _maxLevel + 1
)

func ParseLevel(text string) (Level, error) { _ = "STUB: not implemented"; return *new(Level), nil }

type leveledEnabler interface {
	LevelEnabler

	Level() Level
}

func LevelOf(enab LevelEnabler) Level { _ = "STUB: not implemented"; return *new(Level) }

func (l Level) String() string { _ = "STUB: not implemented"; return "" }

func (l Level) CapitalString() string { _ = "STUB: not implemented"; return "" }

func (l Level) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (l *Level) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (l *Level) unmarshalText(text []byte) bool { _ = "STUB: not implemented"; return false }

func (l *Level) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (l *Level) Get() interface{} { _ = "STUB: not implemented"; return nil }

func (l Level) Enabled(lvl Level) bool { _ = "STUB: not implemented"; return false }

type LevelEnabler interface {
	Enabled(Level) bool
}
