package buffer

import (
	"time"
)

const _size = 1024

type Buffer struct {
	bs   []byte
	pool Pool
}

func (b *Buffer) AppendByte(v byte) { _ = "STUB: not implemented"; return }

func (b *Buffer) AppendBytes(v []byte) { _ = "STUB: not implemented"; return }

func (b *Buffer) AppendString(s string) { _ = "STUB: not implemented"; return }

func (b *Buffer) AppendInt(i int64) { _ = "STUB: not implemented"; return }

func (b *Buffer) AppendTime(t time.Time, layout string) { _ = "STUB: not implemented"; return }

func (b *Buffer) AppendUint(i uint64) { _ = "STUB: not implemented"; return }

func (b *Buffer) AppendBool(v bool) { _ = "STUB: not implemented"; return }

func (b *Buffer) AppendFloat(f float64, bitSize int) { _ = "STUB: not implemented"; return }

func (b *Buffer) Len() int { _ = "STUB: not implemented"; return 0 }

func (b *Buffer) Cap() int { _ = "STUB: not implemented"; return 0 }

func (b *Buffer) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (b *Buffer) String() string { _ = "STUB: not implemented"; return "" }

func (b *Buffer) Reset() { _ = "STUB: not implemented"; return }

func (b *Buffer) Write(bs []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *Buffer) WriteByte(v byte) error { _ = "STUB: not implemented"; return nil }

func (b *Buffer) WriteString(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *Buffer) TrimNewline() { _ = "STUB: not implemented"; return }

func (b *Buffer) Free() { _ = "STUB: not implemented"; return }
