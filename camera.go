package aravis

// #cgo pkg-config: aravis-0.6
// #include <arv.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

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

//export streamCallback
func streamCallback() {
	fmt.Println("streamCallback()")
}

func NewCamera(name string) (Camera, error) {
	var c Camera
	var err error

	cs := C.CString(name)
	c.camera, err = C.arv_camera_new(cs)
	C.free(unsafe.Pointer(cs))

	return c, err
}

func (c *Camera) CreateStream() (Stream, error) {
	var s Stream
	var err error

	s.stream, err = C.arv_camera_create_stream(
		c.camera,
		//		C.stream_callback((C.callback_fcn)(unsafe.Pointer(C.),
		nil,
		nil,
	)

	return s, err
}

func (c *Camera) GetDevice() (Device, error) {
	var d Device
	var err error

	d.device, err = C.arv_camera_get_device(c.camera)

	return d, err
}

func (c *Camera) GetVendorName() (string, error) {
	name, error := C.arv_camera_get_vendor_name(c.camera)
	return C.GoString(name), error
}

func (c *Camera) GetModelName() (string, error) {
	name, error := C.arv_camera_get_model_name(c.camera)
	return C.GoString(name), error
}

func (c *Camera) GetDeviceId() (string, error) {
	id, error := C.arv_camera_get_device_id(c.camera)
	return C.GoString(id), error
}

func (c *Camera) GetSensorSize() (int, int, error) {
	var width, height int
	_, err := C.arv_camera_get_sensor_size(
		c.camera,
		(*C.int)(unsafe.Pointer(&width)),
		(*C.int)(unsafe.Pointer(&height)),
	)
	return int(width), int(height), err
}

func (c *Camera) SetRegion(x, y, width, height int) {
	C.arv_camera_set_region(c.camera,
		C.int(x),
		C.int(y),
		C.int(width),
		C.int(height),
	)
}

func (c *Camera) GetRegion() (int, int, int, int, error) {
	var x, y, width, height int
	_, err := C.arv_camera_get_region(
		c.camera,
		(*C.int)(unsafe.Pointer(&x)),
		(*C.int)(unsafe.Pointer(&y)),
		(*C.int)(unsafe.Pointer(&width)),
		(*C.int)(unsafe.Pointer(&height)),
	)
	return int(x), int(y), int(width), int(height), err
}

func (c *Camera) GetHeightBounds() (int, int, error) {
	var min, max int
	_, err := C.arv_camera_get_height_bounds(
		c.camera,
		(*C.int)(unsafe.Pointer(&min)),
		(*C.int)(unsafe.Pointer(&max)),
	)
	return int(min), int(max), err
}

func (c *Camera) GetWidthBounds() (int, int, error) {
	var min, max int
	_, err := C.arv_camera_get_width_bounds(
		c.camera,
		(*C.int)(unsafe.Pointer(&min)),
		(*C.int)(unsafe.Pointer(&max)),
	)
	return int(min), int(max), err
}

func (c *Camera) SetBinning() {
	// TODO
}

func (c *Camera) GetBinning() (int, int, error) {
	var min, max int
	_, err := C.arv_camera_get_binning(
		c.camera,
		(*C.int)(unsafe.Pointer(&min)),
		(*C.int)(unsafe.Pointer(&max)),
	)
	return int(min), int(max), err
}

func (c *Camera) SetPixelFormat() {
	// TODO
}

func (c *Camera) GetPixelFormat() {
	// TODO
}

func (c *Camera) GetPixelFormatAsString() {
	// TODO
}

func (c *Camera) SetPixelFormatFromString() {
	// TODO
}

func (c *Camera) GetAvailablePixelFormats() {
	// TODO
}

func (c *Camera) GetAvailablePixelFormatsAsDisplayNames() {
	// TODO
}

func (c *Camera) GetAvailablePixelFormatsAsStrings() {
	// TODO
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

func (c *Camera) GVGetCurrentStreamChannel() (int, error) {
	cint, err := C.arv_camera_gv_get_current_stream_channel(c.camera)
	return int(cint), err
}

func (c *Camera) GVGetPacketDelay() (int64, error) {
	cint64, err := C.arv_camera_gv_get_packet_delay(c.camera)
	return int64(cint64), err
}

func (c *Camera) GVSetPacketDelay(delay int64) {
	C.arv_camera_gv_set_packet_delay(c.camera, C.long(delay))
}

func (c *Camera) GVGetPacketSize() (int, error) {
	csize, err := C.arv_camera_gv_get_packet_size(c.camera)
	return int(csize), err
}

func (c *Camera) GVSetPacketSize(size int) {
	C.arv_camera_gv_set_packet_size(c.camera, C.int(size))
}

func (c *Camera) GetChunkMode() (bool, error) {
	mode, err := C.arv_camera_get_chunk_mode(c.camera)
	return toBool(mode), err
}
