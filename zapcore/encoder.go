package zapcore

import (
	"io"
	"time"

	"go.uber.org/zap/buffer"
)

const DefaultLineEnding = "\n"

const OmitKey = ""

type LevelEncoder func(Level, PrimitiveArrayEncoder)

func LowercaseLevelEncoder(l Level, enc PrimitiveArrayEncoder) { _ = "STUB: not implemented"; return }

func LowercaseColorLevelEncoder(l Level, enc PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

func CapitalLevelEncoder(l Level, enc PrimitiveArrayEncoder) { _ = "STUB: not implemented"; return }

func CapitalColorLevelEncoder(l Level, enc PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

func (e *LevelEncoder) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

type TimeEncoder func(time.Time, PrimitiveArrayEncoder)

func EpochTimeEncoder(t time.Time, enc PrimitiveArrayEncoder) { _ = "STUB: not implemented"; return }

func EpochMillisTimeEncoder(t time.Time, enc PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

func EpochNanosTimeEncoder(t time.Time, enc PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

func encodeTimeLayout(t time.Time, layout string, enc PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

func ISO8601TimeEncoder(t time.Time, enc PrimitiveArrayEncoder) { _ = "STUB: not implemented"; return }

func RFC3339TimeEncoder(t time.Time, enc PrimitiveArrayEncoder) { _ = "STUB: not implemented"; return }

func RFC3339NanoTimeEncoder(t time.Time, enc PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

func TimeEncoderOfLayout(layout string) TimeEncoder {
	_ = "STUB: not implemented"
	return *new(TimeEncoder)
}

func (e *TimeEncoder) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (e *TimeEncoder) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *TimeEncoder) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type DurationEncoder func(time.Duration, PrimitiveArrayEncoder)

func SecondsDurationEncoder(d time.Duration, enc PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

func NanosDurationEncoder(d time.Duration, enc PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

func MillisDurationEncoder(d time.Duration, enc PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

func StringDurationEncoder(d time.Duration, enc PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

func (e *DurationEncoder) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

type CallerEncoder func(EntryCaller, PrimitiveArrayEncoder)

func FullCallerEncoder(caller EntryCaller, enc PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

func ShortCallerEncoder(caller EntryCaller, enc PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

func (e *CallerEncoder) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

type NameEncoder func(string, PrimitiveArrayEncoder)

func FullNameEncoder(loggerName string, enc PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

func (e *NameEncoder) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

type EncoderConfig struct {
	MessageKey     string `json:"messageKey" yaml:"messageKey"`
	LevelKey       string `json:"levelKey" yaml:"levelKey"`
	TimeKey        string `json:"timeKey" yaml:"timeKey"`
	NameKey        string `json:"nameKey" yaml:"nameKey"`
	CallerKey      string `json:"callerKey" yaml:"callerKey"`
	FunctionKey    string `json:"functionKey" yaml:"functionKey"`
	StacktraceKey  string `json:"stacktraceKey" yaml:"stacktraceKey"`
	SkipLineEnding bool   `json:"skipLineEnding" yaml:"skipLineEnding"`
	LineEnding     string `json:"lineEnding" yaml:"lineEnding"`

	EncodeLevel    LevelEncoder    `json:"levelEncoder" yaml:"levelEncoder"`
	EncodeTime     TimeEncoder     `json:"timeEncoder" yaml:"timeEncoder"`
	EncodeDuration DurationEncoder `json:"durationEncoder" yaml:"durationEncoder"`
	EncodeCaller   CallerEncoder   `json:"callerEncoder" yaml:"callerEncoder"`

	EncodeName NameEncoder `json:"nameEncoder" yaml:"nameEncoder"`

	NewReflectedEncoder func(io.Writer) ReflectedEncoder `json:"-" yaml:"-"`

	ConsoleSeparator string `json:"consoleSeparator" yaml:"consoleSeparator"`
}

type ObjectEncoder interface {
	AddArray(key string, marshaler ArrayMarshaler) error
	AddObject(key string, marshaler ObjectMarshaler) error

	AddBinary(key string, value []byte)
	AddByteString(key string, value []byte)
	AddBool(key string, value bool)
	AddComplex128(key string, value complex128)
	AddComplex64(key string, value complex64)
	AddDuration(key string, value time.Duration)
	AddFloat64(key string, value float64)
	AddFloat32(key string, value float32)
	AddInt(key string, value int)
	AddInt64(key string, value int64)
	AddInt32(key string, value int32)
	AddInt16(key string, value int16)
	AddInt8(key string, value int8)
	AddString(key, value string)
	AddTime(key string, value time.Time)
	AddUint(key string, value uint)
	AddUint64(key string, value uint64)
	AddUint32(key string, value uint32)
	AddUint16(key string, value uint16)
	AddUint8(key string, value uint8)
	AddUintptr(key string, value uintptr)

	AddReflected(key string, value interface{}) error

	OpenNamespace(key string)
}

type ArrayEncoder interface {
	PrimitiveArrayEncoder

	AppendDuration(time.Duration)
	AppendTime(time.Time)

	AppendArray(ArrayMarshaler) error
	AppendObject(ObjectMarshaler) error

	AppendReflected(value interface{}) error
}

type PrimitiveArrayEncoder interface {
	AppendBool(bool)
	AppendByteString([]byte)
	AppendComplex128(complex128)
	AppendComplex64(complex64)
	AppendFloat64(float64)
	AppendFloat32(float32)
	AppendInt(int)
	AppendInt64(int64)
	AppendInt32(int32)
	AppendInt16(int16)
	AppendInt8(int8)
	AppendString(string)
	AppendUint(uint)
	AppendUint64(uint64)
	AppendUint32(uint32)
	AppendUint16(uint16)
	AppendUint8(uint8)
	AppendUintptr(uintptr)
}

type Encoder interface {
	ObjectEncoder

	Clone() Encoder

	EncodeEntry(Entry, []Field) (*buffer.Buffer, error)
}
