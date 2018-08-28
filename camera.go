package aravis

// #cgo pkg-config: aravis-0.6
// #include <arv.h>
// #include <stdlib.h>
import "C"
import "unsafe"

type Camera struct {
	camera *C.struct__ArvCamera
}

const (
	ACQUISITION_MODE_CONTINUOUS   = C.ARV_ACQUISITION_MODE_CONTINUOUS
	ACQUISITION_MODE_SINGLE_FRAME = C.ARV_ACQUISITION_MODE_SINGLE_FRAME
)

const (
	ARV_AUTO_OFF        = C.ARV_AUTO_OFF
	ARV_AUTO_ONCE       = C.ARV_AUTO_ONCE
	ARV_AUTO_CONTINUOUS = C.ARV_AUTO_CONTINUOUS
)

func NewCamera(name string) (Camera, error) {
	var err error
	var c Camera

	cs := C.CString(name)
	c.camera, err = C.arv_camera_new(cs)
	C.free(unsafe.Pointer(cs))

	return c, err
}

func (c *Camera) StartAcquisition() {
	C.arv_camera_start_acquisition(c.camera)
}

func (c *Camera) StopAcquisition() {
	C.arv_camera_stop_acquisition(c.camera)
}

func (c *Camera) AbortAcquisition() {
	C.arv_camera_abort_acquisition(c.camera)
}

func (c *Camera) SetFrameRate(frameRate float64) {
	C.arv_camera_set_frame_rate(c.camera, C.double(frameRate))
}

func (c *Camera) GetFrameRate() (float64, error) {
	fr, err := C.arv_camera_get_frame_rate(c.camera)
	return float64(fr), err
}

func (c *Camera) GetFrameRateBounds() (float64, float64, error) {
	var min, max float64
	_, err := C.arv_camera_get_frame_rate_bounds(
		c.camera,
		(*C.double)(unsafe.Pointer(&min)),
		(*C.double)(unsafe.Pointer(&max)),
	)
	return float64(min), float64(max), err
}

func (c *Camera) SetTrigger(source string) {
	csource := C.CString(source)
	C.arv_camera_set_trigger(c.camera, csource)
	C.free(unsafe.Pointer(csource))
}

func (c *Camera) SetTriggerSource(source string) {
	csource := C.CString(source)
	C.arv_camera_set_trigger_source(c.camera, csource)
	C.free(unsafe.Pointer(csource))
}

func (c *Camera) GetTriggerSource() (string, error) {
	csource, err := C.arv_camera_get_trigger_source(c.camera)
	return C.GoString(csource), err
}

func (c *Camera) SoftwareTrigger() {
	C.arv_camera_software_trigger(c.camera)
}

func (c *Camera) SetGain(gain float64) {
	C.arv_camera_set_gain(c.camera, C.double(gain))
}

func (c *Camera) GetGain() (float64, error) {
	cgain, err := C.arv_camera_get_gain(c.camera)
	return float64(cgain), err
}

func (c *Camera) GetGainBounds() (float64, float64, error) {
	var min, max float64
	_, err := C.arv_camera_get_gain_bounds(
		c.camera,
		(*C.double)(unsafe.Pointer(&min)),
		(*C.double)(unsafe.Pointer(&max)),
	)
	return float64(min), float64(max), err
}

func (c *Camera) SetGainAuto() {
	// TODO
}

func (c *Camera) GetPayloadSize() (uint, error) {
	csize, err := C.arv_camera_get_payload(c.camera)
	return uint(csize), err
}

func (c *Camera) IsGVDevice() (bool, error) {
	cbool, err := C.arv_camera_is_gv_device(c.camera)
	return toBool(cbool), err
}

func (c *Camera) GVGetNumStreamChannels() (int, error) {
	cint, err := C.arv_camera_gv_get_n_stream_channels(c.camera)
	return int(cint), err
}

func (c *Camera) GVSelectStreamChannels(id int) {
	C.arv_camera_gv_select_stream_channel(c.camera, C.gint(id))
}
