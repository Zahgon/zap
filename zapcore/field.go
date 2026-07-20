package zapcore

type FieldType uint8

const (
	UnknownType FieldType = iota

	ArrayMarshalerType

	ObjectMarshalerType

	BinaryType

	BoolType

	ByteStringType

	Complex128Type

	Complex64Type

	DurationType

	Float64Type

	Float32Type

	Int64Type

	Int32Type

	Int16Type

	Int8Type

	StringType

	TimeType

	TimeFullType

	Uint64Type

	Uint32Type

	Uint16Type

	Uint8Type

	UintptrType

	ReflectType

	NamespaceType

	StringerType

	ErrorType

	SkipType

	InlineMarshalerType
)

type Field struct {
	Key       string
	Type      FieldType
	Integer   int64
	String    string
	Interface interface{}
}

func (f Field) AddTo(enc ObjectEncoder) { _ = "STUB: not implemented"; return }

func (f Field) Equals(other Field) bool { _ = "STUB: not implemented"; return false }

func addFields(enc ObjectEncoder, fields []Field) { _ = "STUB: not implemented"; return }

func encodeStringer(key string, stringer interface{}, enc ObjectEncoder) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}
