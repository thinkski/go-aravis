package aravis

// #cgo pkg-config: aravis-0.6
// #include <arv.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
	"time"
)

type Stream struct {
	stream *C.struct__ArvStream
}

func (s *Stream) PushBuffer(b Buffer) {
	C.arv_stream_push_buffer(s.stream, b.buffer)
}

func (s *Stream) PopBuffer() (Buffer, error) {
	var b Buffer
	var err error

	b.buffer, err = C.arv_stream_pop_buffer(s.stream)

	return b, err
}

func (s *Stream) TryPopBuffer() (Buffer, error) {
	var b Buffer
	var err error

	b.buffer, err = C.arv_stream_try_pop_buffer(s.stream)

	return b, err
}

func (s *Stream) TimeoutPopBuffer(t time.Duration) (Buffer, error) {
	var b Buffer
	var err error

	b.buffer, err = C.arv_stream_timeout_pop_buffer(s.stream, C.guint64(t/1000))

	if b.buffer == nil {
		return Buffer{}, errors.New("Aravis returned null pointer")
	}

	return b, err
}

func (s *Stream) Close() {
	C.g_object_unref(C.gpointer(s.stream))
}
