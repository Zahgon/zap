package zapcore

type ObjectMarshaler interface {
	MarshalLogObject(ObjectEncoder) error
}

type ObjectMarshalerFunc func(ObjectEncoder) error

func (f ObjectMarshalerFunc) MarshalLogObject(enc ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type ArrayMarshaler interface {
	MarshalLogArray(ArrayEncoder) error
}

type ArrayMarshalerFunc func(ArrayEncoder) error

func (f ArrayMarshalerFunc) MarshalLogArray(enc ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}
