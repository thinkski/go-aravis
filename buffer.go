package aravis

// #cgo pkg-config: aravis-0.6
// #include <arv.h>
import "C"

type Buffer struct {
	buffer *C.struct__ArvBuffer
}

func NewBuffer(size uint) (Buffer, error) {
	var b Buffer
	var err error

	b.buffer, err = C.arv_buffer_new(C.size_t(size), nil)

	return b, err
}
