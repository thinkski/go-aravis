package aravis

// #cgo pkg-config: aravis-0.6
// #include <arv.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func GetDeviceId(index uint) (string, error) {
	s, err := C.arv_get_device_id(C.uint(index))
	return C.GoString(s), err
}

func GetInterfaceId(index uint) (string, error) {
	s, err := C.arv_get_interface_id(C.uint(index))
	return C.GoString(s), err
}

func DisableInterface(id string) {
	cs := C.CString(id)
	C.arv_disable_interface(cs)
	C.free(unsafe.Pointer(cs))
}

func EnableInterface(id string) {
	cs := C.CString(id)
	C.arv_enable_interface(cs)
	C.free(unsafe.Pointer(cs))
}

func GetNumDevices() (uint, error) {
	n, err := C.arv_get_n_devices()
	return uint(n), err
}

func GetNumInferface() (uint, error) {
	n, err := C.arv_get_n_interfaces()
	return uint(n), err
}

func UpdateDeviceList() {
	C.arv_update_device_list()
}

func OpenDevice() {
	// TODO
}

func Shutdown() {
	C.arv_shutdown()
}

func InterfaceGetDeviceId() {
	// TODO
}

func InterfaceGetDevicePhysicalId() {
	// TODO
}

func InterfaceGetDeviceAddress() {
	// TODO
}

func InterfaceGetNumDevices() {
	// TODO
}

func InterfaceOpenDevice() {
	// TODO
}
