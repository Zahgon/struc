package struc

import (
	"io"
)

type Float16 float64

func (f *Float16) Pack(p []byte, opt *Options) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *Float16) Unpack(r io.Reader, length int, opt *Options) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *Float16) Size(opt *Options) int { _ = "STUB: not implemented"; return 0 }

func (f *Float16) String() string { _ = "STUB: not implemented"; return "" }
