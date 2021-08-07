package sony

import "fmt"

type SonyTV struct {
	vol     int
	channel int
	isOn    bool
}

func NewTV(vol int, channel int, isOn bool) *SonyTV {

	return &SonyTV{vol: vol, channel: channel, isOn: isOn}
}

func (st *SonyTV) TurnOn() {
	fmt.Println("SonyTV is now on")
	st.isOn = true
}
func (st *SonyTV) TurnOff() {
	fmt.Println("SonyTV is now off")
	st.isOn = false
}

func (st *SonyTV) VolumeUp() int {
	st.vol++
	fmt.Println("Increasing SonyTV volume to", st.vol)
	return st.vol
}
func (st *SonyTV) VolumeDown() int {
	st.vol--
	fmt.Println("Decreasing SonyTV volume to", st.vol)
	return st.vol
}

func (st *SonyTV) ChannelUp() int {
	st.channel++
	fmt.Println("Decreasing SonyTV channel to", st.channel)
	return st.channel
}
func (st *SonyTV) ChannelDown() int {
	st.channel--
	fmt.Println("Decreasing SonyTV channel to", st.channel)
	return st.channel
}

func (st *SonyTV) GoToChannel(ch int) {
	st.channel = ch
	fmt.Println("Setting SonyTV channel to", st.channel)
}
