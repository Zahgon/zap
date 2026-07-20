package zapio

import (
	"bytes"
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Writer struct {
	Log *zap.Logger

	Level zapcore.Level

	buff bytes.Buffer
}

var (
	_ zapcore.WriteSyncer = (*Writer)(nil)
	_ io.Closer           = (*Writer)(nil)
)

func (w *Writer) Write(bs []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (w *Writer) writeLine(line []byte) (remaining []byte) { _ = "STUB: not implemented"; return nil }

func (w *Writer) Close() error { _ = "STUB: not implemented"; return nil }

func (w *Writer) Sync() error { _ = "STUB: not implemented"; return nil }

func (w *Writer) flush(allowEmpty bool) { _ = "STUB: not implemented"; return }

func (w *Writer) log(b []byte) { _ = "STUB: not implemented"; return }
