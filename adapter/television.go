package main

// This is the television interface we want to use with both TV types
type television interface {
	TurnOn()
	TurnOff()
	ChannelUp() int
	ChannelDown() int
	GoToChannel(ch int)
	VolumeUp() int
	VolumeDown() int
}
