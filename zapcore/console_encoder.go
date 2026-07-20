package zapcore

import (
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/internal/pool"
)

var _sliceEncoderPool = pool.New(func() *sliceArrayEncoder {
	return &sliceArrayEncoder{
		elems: make([]interface{}, 0, 2),
	}
})

func getSliceEncoder() *sliceArrayEncoder { _ = "STUB: not implemented"; return nil }

func putSliceEncoder(e *sliceArrayEncoder) { _ = "STUB: not implemented"; return }

type consoleEncoder struct {
	*jsonEncoder
}

func NewConsoleEncoder(cfg EncoderConfig) Encoder { _ = "STUB: not implemented"; return *new(Encoder) }

func (c consoleEncoder) Clone() Encoder { _ = "STUB: not implemented"; return *new(Encoder) }

func (c consoleEncoder) EncodeEntry(ent Entry, fields []Field) (*buffer.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c consoleEncoder) writeContext(line *buffer.Buffer, extra []Field) {
	_ = "STUB: not implemented"
	return
}

func (c consoleEncoder) addSeparatorIfNecessary(line *buffer.Buffer) {
	_ = "STUB: not implemented"
	return
}
