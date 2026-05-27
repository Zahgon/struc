package struc

import (
	"encoding/binary"
	"io"
	"reflect"
)

type Options struct {
	ByteAlign int
	PtrSize   int
	Order     binary.ByteOrder
}

func (o *Options) Validate() error { _ = "STUB: not implemented"; return nil }

var emptyOptions = &Options{}

func init() {
	// fill default values to avoid data race to be reported by race detector.
	emptyOptions.Validate()
}

func prep(data interface{}) (reflect.Value, Packer, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(Packer), nil
}

func Pack(w io.Writer, data interface{}) error { _ = "STUB: not implemented"; return nil }

func PackWithOptions(w io.Writer, data interface{}, options *Options) error {
	_ = "STUB: not implemented"
	return nil
}

func Unpack(r io.Reader, data interface{}) error { _ = "STUB: not implemented"; return nil }

func UnpackWithOptions(r io.Reader, data interface{}, options *Options) error {
	_ = "STUB: not implemented"
	return nil
}

func Sizeof(data interface{}) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func SizeofWithOptions(data interface{}, options *Options) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
