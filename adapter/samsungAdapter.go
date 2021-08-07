package main

import "adapter/samsung"

type samsungAdapter struct {
	// A field for the samsungTV reference
	sstv *samsung.SamsungTV
}

func (ss *samsungAdapter) TurnOn() {
	ss.sstv.SetOnState(true)
}
func (ss *samsungAdapter) TurnOff() {
	ss.sstv.SetOnState(false)
}

func (ss *samsungAdapter) VolumeUp() int {
	vol := ss.sstv.GetVolume() + 1
	ss.sstv.SetVolume(vol)
	return vol
}
func (ss *samsungAdapter) VolumeDown() int {
	vol := ss.sstv.GetVolume() - 1
	ss.sstv.SetVolume(vol)
	return vol
}

func (ss *samsungAdapter) ChannelUp() int {
	ch := ss.sstv.GetChannel() + 1
	ss.sstv.SetChannel(ch)
	return ch
}

func (ss *samsungAdapter) ChannelDown() int {
	ch := ss.sstv.GetChannel() - 1
	ss.sstv.SetChannel(ch)
	return ch
}

func (ss *samsungAdapter) GoToChannel(ch int) {
	ss.sstv.SetChannel(ch)
}
