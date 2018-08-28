package main

import (
	"log"
	"time"

	aravis "github.com/thinkski/go-aravis"
)

func main() {
	var err error
	var numDevices uint

	// Get devices
	aravis.UpdateDeviceList()
	if numDevices, err = aravis.GetNumDevices(); err != nil {
		log.Fatal(err)
	}

	// Must find at least one device
	if numDevices == 0 {
		log.Fatal("No devices found. Exiting.")
		return
	}

	name, err := aravis.GetDeviceId(0)
	log.Println(name)

	camera, err := aravis.NewCamera(name)

	camera.SetFrameRate(10)
	size, err := camera.GetPayloadSize()

	stream, err := camera.CreateStream()

	buffer, err := aravis.NewBuffer(size)

	stream.PushBuffer(&buffer)

	camera.StartAcquisition()

	time.Sleep(3 * time.Second)

	log.Println(size)
}
