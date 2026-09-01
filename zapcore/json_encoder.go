package zapcore

import (
	"time"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/internal/pool"
)

const _hex = "0123456789abcdef"

var _jsonPool = pool.New(func() *jsonEncoder {
	return &jsonEncoder{}
})

func putJSONEncoder(enc *jsonEncoder) { _ = "STUB: not implemented"; return }

type jsonEncoder struct {
	*EncoderConfig
	buf            *buffer.Buffer
	spaced         bool
	openNamespaces int

	reflectBuf *buffer.Buffer
	reflectEnc ReflectedEncoder
}

func NewJSONEncoder(cfg EncoderConfig) Encoder { _ = "STUB: not implemented"; return *new(Encoder) }

func newJSONEncoder(cfg EncoderConfig, spaced bool) *jsonEncoder {
	_ = "STUB: not implemented"
	return nil
}

func (enc *jsonEncoder) AddArray(key string, arr ArrayMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *jsonEncoder) AddObject(key string, obj ObjectMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *jsonEncoder) AddBinary(key string, val []byte) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AddByteString(key string, val []byte) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AddBool(key string, val bool) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AddComplex128(key string, val complex128) {
	_ = "STUB: not implemented"
	return
}

func (enc *jsonEncoder) AddComplex64(key string, val complex64) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AddDuration(key string, val time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (enc *jsonEncoder) AddFloat64(key string, val float64) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AddFloat32(key string, val float32) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AddInt64(key string, val int64) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) resetReflectBuf() { _ = "STUB: not implemented"; return }

var nullLiteralBytes = []byte("null")

func (enc *jsonEncoder) encodeReflected(obj interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (enc *jsonEncoder) AddReflected(key string, obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *jsonEncoder) OpenNamespace(key string) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AddString(key, val string) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AddTime(key string, val time.Time) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AddUint64(key string, val uint64) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AppendArray(arr ArrayMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *jsonEncoder) AppendObject(obj ObjectMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *jsonEncoder) AppendBool(val bool) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AppendByteString(val []byte) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) appendComplex(val complex128, precision int) {
	_ = "STUB: not implemented"
	return
}

func (enc *jsonEncoder) AppendDuration(val time.Duration) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AppendInt64(val int64) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AppendReflected(val interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (enc *jsonEncoder) AppendString(val string) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AppendTimeLayout(time time.Time, layout string) {
	_ = "STUB: not implemented"
	return
}

func (enc *jsonEncoder) AppendTime(val time.Time) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AppendUint64(val uint64) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) AddInt(k string, v int)         { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AddInt32(k string, v int32)     { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AddInt16(k string, v int16)     { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AddInt8(k string, v int8)       { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AddUint(k string, v uint)       { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AddUint32(k string, v uint32)   { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AddUint16(k string, v uint16)   { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AddUint8(k string, v uint8)     { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AddUintptr(k string, v uintptr) { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AppendComplex64(v complex64)    { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AppendComplex128(v complex128)  { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AppendFloat64(v float64)        { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AppendFloat32(v float32)        { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AppendInt(v int)                { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AppendInt32(v int32)            { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AppendInt16(v int16)            { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AppendInt8(v int8)              { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AppendUint(v uint)              { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AppendUint32(v uint32)          { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AppendUint16(v uint16)          { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AppendUint8(v uint8)            { _ = "STUB: not implemented"; return }
func (enc *jsonEncoder) AppendUintptr(v uintptr)        { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) Clone() Encoder { _ = "STUB: not implemented"; return *new(Encoder) }

func (enc *jsonEncoder) clone() *jsonEncoder { _ = "STUB: not implemented"; return nil }

func (enc *jsonEncoder) EncodeEntry(ent Entry, fields []Field) (*buffer.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (enc *jsonEncoder) truncate() { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) closeOpenNamespaces() { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) addKey(key string) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) addElementSeparator() { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) appendFloat(val float64, bitSize int) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) safeAddString(s string) { _ = "STUB: not implemented"; return }

func (enc *jsonEncoder) safeAddByteString(s []byte) { _ = "STUB: not implemented"; return }

func safeAppendStringLike[S []byte | string](

	appendTo func(*buffer.Buffer, S),

	decodeRune func(S) (rune, int),
	buf *buffer.Buffer,
	s S,
) {
	_ = "STUB: not implemented"
	return
}
