package struc

import (
	"io"
	"reflect"
)

type byteWriter struct {
	buf []byte
	pos int
}

func (b byteWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type binaryFallback reflect.Value

func (b binaryFallback) String() string { _ = "STUB: not implemented"; return "" }

func (b binaryFallback) Sizeof(val reflect.Value, options *Options) int {
	_ = "STUB: not implemented"
	return 0
}

func (b binaryFallback) Pack(buf []byte, val reflect.Value, options *Options) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b binaryFallback) Unpack(r io.Reader, val reflect.Value, options *Options) error {
	_ = "STUB: not implemented"
	return nil
}
