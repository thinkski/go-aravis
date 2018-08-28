package aravis

/*
#include <arv.h>

void streamCallback_cgo(
	void *user_data,
	ArvStreamCallbackType type,
	ArvBuffer *buffer
) {
	streamCallback();
	return;
}
*/
import "C"
