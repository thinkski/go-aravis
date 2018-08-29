package aravis

// #cgo pkg-config: aravis-0.6
// #include <arv.h>
import "C"
import "unsafe"

const (
	BUFFER_STATUS_UNKNOWN         = C.ARV_BUFFER_STATUS_UNKNOWN
	BUFFER_STATUS_SUCCESS         = C.ARV_BUFFER_STATUS_SUCCESS
	BUFFER_STATUS_CLEARED         = C.ARV_BUFFER_STATUS_CLEARED
	BUFFER_STATUS_TIMEOUT         = C.ARV_BUFFER_STATUS_TIMEOUT
	BUFFER_STATUS_MISSING_PACKETS = C.ARV_BUFFER_STATUS_MISSING_PACKETS
	BUFFER_STATUS_WRONG_PACKET_ID = C.ARV_BUFFER_STATUS_WRONG_PACKET_ID
	BUFFER_STATUS_SIZE_MISMATCH   = C.ARV_BUFFER_STATUS_SIZE_MISMATCH
	BUFFER_STATUS_FILLING         = C.ARV_BUFFER_STATUS_FILLING
	BUFFER_STATUS_ABORTED         = C.ARV_BUFFER_STATUS_ABORTED
)

type Buffer struct {
	buffer *C.struct__ArvBuffer
}

func NewBuffer(size uint) (Buffer, error) {
	var b Buffer

	if buffer, err := C.arv_buffer_new(C.size_t(size), nil); err != nil || buffer == nil {
		return Buffer{nil}, err
	} else {
		b.buffer = buffer
		return b, err
	}
}

func (b *Buffer) GetData() ([]byte, error) {
	var size int

	data, err := C.arv_buffer_get_data(
		b.buffer,
		(*C.size_t)(unsafe.Pointer(&size)),
	)

	return C.GoBytes(data, C.int(size)), err
}

func (b *Buffer) GetStatus() (int, error) {
	status, err := C.arv_buffer_get_status(b.buffer)
	return int(status), err
}
