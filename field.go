package struc

import (
	"encoding/binary"
	"reflect"
)

type Field struct {
	Name     string
	Ptr      bool
	Index    int
	Type     Type
	defType  Type
	Array    bool
	Slice    bool
	Len      int
	Order    binary.ByteOrder
	Sizeof   []int
	Sizefrom []int
	Fields   Fields
	kind     reflect.Kind
}

func (f *Field) String() string { _ = "STUB: not implemented"; return "" }

func (f *Field) Size(val reflect.Value, options *Options) int { _ = "STUB: not implemented"; return 0 }

func (f *Field) packVal(buf []byte, val reflect.Value, length int, options *Options) (size int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TODO: handle kind != bytes here

func (f *Field) Pack(buf []byte, val reflect.Value, length int, options *Options) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// special case strings and byte slices for performance

// TODO: allow configuring pad byte?

func (f *Field) unpackVal(buf []byte, val reflect.Value, length int, options *Options) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Field) Unpack(buf []byte, val reflect.Value, length int, options *Options) error {
	_ = "STUB: not implemented"
	return nil
}

// special case byte slices for performance
