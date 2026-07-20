package zapcore

import "sync"

type lazyWithCore struct {
	core         Core
	originalCore Core
	sync.Once
	fields []Field
}

func NewLazyWith(core Core, fields []Field) Core { _ = "STUB: not implemented"; return *new(Core) }

func (d *lazyWithCore) initOnce() { _ = "STUB: not implemented"; return }

func (d *lazyWithCore) With(fields []Field) Core { _ = "STUB: not implemented"; return *new(Core) }

func (d *lazyWithCore) Check(e Entry, ce *CheckedEntry) *CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (d *lazyWithCore) Enabled(level Level) bool { _ = "STUB: not implemented"; return false }

func (d *lazyWithCore) Write(e Entry, fields []Field) error { _ = "STUB: not implemented"; return nil }

func (d *lazyWithCore) Sync() error { _ = "STUB: not implemented"; return nil }
