![Go Report Card](https://goreportcard.com/badge/github.com/thinkski/go-aravis)
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

## Troubleshooting

GigE Vision cameras often use large packet sizes, well in excess of the typical MTU (maximum transmit unit) of most network interface cards, to save on packet overhead. Be sure to first set the MTU of the network interface(s) with GigE Vision cameras to 9000 bytes. For instance, if the network interface is `enp2s0`:

    ip link set enp2s0 mtu 9000
