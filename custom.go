package struc

import (
	"io"
	"reflect"
)

type Custom interface {
	Pack(p []byte, opt *Options) (int, error)
	Unpack(r io.Reader, length int, opt *Options) error
	Size(opt *Options) int
	String() string
}

type customFallback struct {
	custom Custom
}

func (c customFallback) Pack(p []byte, val reflect.Value, opt *Options) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c customFallback) Unpack(r io.Reader, val reflect.Value, opt *Options) error {
	_ = "STUB: not implemented"
	return nil
}

func (c customFallback) Sizeof(val reflect.Value, opt *Options) int {
	_ = "STUB: not implemented"
	return 0
}

func (c customFallback) String() string { _ = "STUB: not implemented"; return "" }
