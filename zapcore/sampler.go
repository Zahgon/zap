package zapcore

import (
	"sync/atomic"
	"time"
)

const (
	_numLevels        = _maxLevel - _minLevel + 1
	_countersPerLevel = 4096
)

type counter struct {
	resetAt atomic.Int64
	counter atomic.Uint64
}

type counters [_numLevels][_countersPerLevel]counter

func newCounters() *counters { _ = "STUB: not implemented"; return nil }

func (cs *counters) get(lvl Level, key string) *counter { _ = "STUB: not implemented"; return nil }

func fnv32a(s string) uint32 { _ = "STUB: not implemented"; return 0 }

func (c *counter) IncCheckReset(t time.Time, tick time.Duration) uint64 {
	_ = "STUB: not implemented"
	return 0
}

type SamplingDecision uint32

const (
	LogDropped SamplingDecision = 1 << iota

	LogSampled
)

type optionFunc func(*sampler)

func (f optionFunc) apply(s *sampler) { _ = "STUB: not implemented"; return }

type SamplerOption interface {
	apply(*sampler)
}

func nopSamplingHook(Entry, SamplingDecision) { _ = "STUB: not implemented"; return }

func SamplerHook(hook func(entry Entry, dec SamplingDecision)) SamplerOption {
	_ = "STUB: not implemented"
	return *new(SamplerOption)
}

func NewSamplerWithOptions(core Core, tick time.Duration, first, thereafter int, opts ...SamplerOption) Core {
	_ = "STUB: not implemented"
	return *new(Core)
}

type sampler struct {
	Core

	counts            *counters
	tick              time.Duration
	first, thereafter uint64
	hook              func(Entry, SamplingDecision)
}

var (
	_ Core           = (*sampler)(nil)
	_ leveledEnabler = (*sampler)(nil)
)

func NewSampler(core Core, tick time.Duration, first, thereafter int) Core {
	_ = "STUB: not implemented"
	return *new(Core)
}

func (s *sampler) Level() Level { _ = "STUB: not implemented"; return *new(Level) }

func (s *sampler) With(fields []Field) Core { _ = "STUB: not implemented"; return *new(Core) }

func (s *sampler) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}
