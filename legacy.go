package struc

import (
	"encoding/binary"
	"io"
)

// Deprecated. Use PackWithOptions.
func PackWithOrder(w io.Writer, data interface{}, order binary.ByteOrder) error {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated. Use UnpackWithOptions.
func UnpackWithOrder(r io.Reader, data interface{}, order binary.ByteOrder) error {
	_ = "STUB: not implemented"
	return nil
}
