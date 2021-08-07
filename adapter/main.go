package main

import (
	"adapter/samsung"
	"adapter/sony"
	"fmt"
)

func main() {

	sonyTV := sony.NewTV(20, 9, true)

	// Because the SonyTV implements the "television" interface, we don't need an adapter
	performTestTV(sonyTV)

	samsTV := samsung.NewTV(13, 35, true)

	// We need to create a SamsungTV adapter for the SamsungTV class, however
	// because it has an interface that's different from the one we want to use
	samsAdapt := &samsungAdapter{sstv: samsTV}
	performTestTV(samsAdapt)
}

func performTestTV(tv television) {
	tv.TurnOn()
	tv.ChannelUp()
	tv.ChannelDown()
	tv.GoToChannel(98)
	tv.VolumeUp()
	tv.VolumeDown()
	tv.TurnOff()
	fmt.Println("--------------------")
}
