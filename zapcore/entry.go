package zapcore

import (
	"time"

	"go.uber.org/zap/internal/pool"
)

var _cePool = pool.New(func() *CheckedEntry {

	return &CheckedEntry{
		cores: make([]Core, 4),
	}
})

func getCheckedEntry() *CheckedEntry { _ = "STUB: not implemented"; return nil }

func putCheckedEntry(ce *CheckedEntry) { _ = "STUB: not implemented"; return }

func NewEntryCaller(pc uintptr, file string, line int, ok bool) EntryCaller {
	_ = "STUB: not implemented"
	return *new(EntryCaller)
}

type EntryCaller struct {
	Defined  bool
	PC       uintptr
	File     string
	Line     int
	Function string
}

func (ec EntryCaller) String() string { _ = "STUB: not implemented"; return "" }

func (ec EntryCaller) FullPath() string { _ = "STUB: not implemented"; return "" }

func (ec EntryCaller) TrimmedPath() string { _ = "STUB: not implemented"; return "" }

type Entry struct {
	Level      Level
	Time       time.Time
	LoggerName string
	Message    string
	Caller     EntryCaller
	Stack      string
}

type CheckWriteHook interface {
	OnWrite(*CheckedEntry, []Field)
}

type CheckWriteAction uint8

const (
	WriteThenNoop CheckWriteAction = iota

	WriteThenGoexit

	WriteThenPanic

	WriteThenFatal
)

func (a CheckWriteAction) OnWrite(ce *CheckedEntry, _ []Field) { _ = "STUB: not implemented"; return }

var _ CheckWriteHook = CheckWriteAction(0)

type CheckPreWriteHook func(Entry, []Field) (Entry, []Field)

type CheckedEntry struct {
	Entry
	ErrorOutput WriteSyncer
	dirty       bool
	after       CheckWriteHook
	cores       []Core
	before      []CheckPreWriteHook
}

func (ce *CheckedEntry) reset() { _ = "STUB: not implemented"; return }

func (ce *CheckedEntry) Write(fields ...Field) { _ = "STUB: not implemented"; return }

func (ce *CheckedEntry) AddCore(ent Entry, core Core) *CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (ce *CheckedEntry) Should(ent Entry, should CheckWriteAction) *CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (ce *CheckedEntry) Before(ent Entry, hook CheckPreWriteHook) *CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (ce *CheckedEntry) After(ent Entry, hook CheckWriteHook) *CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}
