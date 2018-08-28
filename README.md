# go-aravis

Go wrapper around libaravis


## Quickstart

How to get the number of connected devices:

    import aravis
    import log

    func main() {
        aravis.UpdateDeviceList()

        n, err := aravis.GetNumDevices()
        if err != nil {
            log.Fatal(err)
        }

        log.Println("Devices:", n)
    }
