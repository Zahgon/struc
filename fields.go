package struc

import (
	"encoding/binary"
	"io"
	"reflect"
)

type Fields []*Field

func (f Fields) SetByteOrder(order binary.ByteOrder) { _ = "STUB: not implemented"; return }

func (f Fields) String() string { _ = "STUB: not implemented"; return "" }

func (f Fields) Sizeof(val reflect.Value, options *Options) int {
	_ = "STUB: not implemented"
	return 0
}

func (f Fields) sizefrom(val reflect.Value, index []int) int { _ = "STUB: not implemented"; return 0 }

// all the builtin array length types are native int
// so this guards against weird truncation

func (f Fields) Pack(buf []byte, val reflect.Value, options *Options) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// allocating a new int here has fewer side effects (doesn't update the original struct)
// but it's a wasteful allocation
// the old method might work if we just cast the temporary int/uint to the target type

func (f Fields) Unpack(r io.Reader, val reflect.Value, options *Options) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: DRY (we repeat the inner loop above)
