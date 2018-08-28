package aravis

// #cgo pkg-config: aravis-0.6
// #include <arv.h>
// #include <stdlib.h>
import "C"

type Stream struct {
	stream *C.struct__ArvStream
}
