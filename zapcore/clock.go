package zapcore

import "time"

var DefaultClock = systemClock{}

type Clock interface {
	Now() time.Time

	NewTicker(time.Duration) *time.Ticker
}

type systemClock struct{}

func (systemClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (systemClock) NewTicker(duration time.Duration) *time.Ticker {
	_ = "STUB: not implemented"
	return nil
}
