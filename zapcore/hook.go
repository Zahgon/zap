package zapcore

type hooked struct {
	Core
	funcs []func(Entry) error
}

var (
	_ Core           = (*hooked)(nil)
	_ leveledEnabler = (*hooked)(nil)
)

func RegisterHooks(core Core, hooks ...func(Entry) error) Core {
	_ = "STUB: not implemented"
	return *new(Core)
}

func (h *hooked) Level() Level { _ = "STUB: not implemented"; return *new(Level) }

func (h *hooked) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (h *hooked) With(fields []Field) Core { _ = "STUB: not implemented"; return *new(Core) }

func (h *hooked) Write(ent Entry, _ []Field) error { _ = "STUB: not implemented"; return nil }
