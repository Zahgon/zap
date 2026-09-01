package zapcore

import "time"

type MapObjectEncoder struct {
	Fields map[string]interface{}

	cur map[string]interface{}
}

func NewMapObjectEncoder() *MapObjectEncoder { _ = "STUB: not implemented"; return nil }

func (m *MapObjectEncoder) AddArray(key string, v ArrayMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MapObjectEncoder) AddObject(k string, v ObjectMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MapObjectEncoder) AddBinary(k string, v []byte) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddByteString(k string, v []byte) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddBool(k string, v bool) { _ = "STUB: not implemented"; return }

func (m MapObjectEncoder) AddDuration(k string, v time.Duration) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddComplex128(k string, v complex128) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddComplex64(k string, v complex64) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddFloat64(k string, v float64) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddFloat32(k string, v float32) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddInt(k string, v int) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddInt64(k string, v int64) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddInt32(k string, v int32) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddInt16(k string, v int16) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddInt8(k string, v int8) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddString(k string, v string) { _ = "STUB: not implemented"; return }

func (m MapObjectEncoder) AddTime(k string, v time.Time) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddUint(k string, v uint) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddUint64(k string, v uint64) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddUint32(k string, v uint32) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddUint16(k string, v uint16) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddUint8(k string, v uint8) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddUintptr(k string, v uintptr) { _ = "STUB: not implemented"; return }

func (m *MapObjectEncoder) AddReflected(k string, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MapObjectEncoder) OpenNamespace(k string) { _ = "STUB: not implemented"; return }

type sliceArrayEncoder struct {
	elems []interface{}
}

func (s *sliceArrayEncoder) AppendArray(v ArrayMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sliceArrayEncoder) AppendObject(v ObjectMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sliceArrayEncoder) AppendReflected(v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sliceArrayEncoder) AppendBool(v bool)              { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendByteString(v []byte)      { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendComplex128(v complex128)  { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendComplex64(v complex64)    { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendDuration(v time.Duration) { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendFloat64(v float64)        { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendFloat32(v float32)        { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendInt(v int)                { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendInt64(v int64)            { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendInt32(v int32)            { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendInt16(v int16)            { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendInt8(v int8)              { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendString(v string)          { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendTime(v time.Time)         { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendUint(v uint)              { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendUint64(v uint64)          { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendUint32(v uint32)          { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendUint16(v uint16)          { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendUint8(v uint8)            { _ = "STUB: not implemented"; return }
func (s *sliceArrayEncoder) AppendUintptr(v uintptr)        { _ = "STUB: not implemented"; return }
