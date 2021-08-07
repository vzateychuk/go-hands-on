package samsung

import "fmt"

type SamsungTV struct {
	currentChan   int
	currentVolume int
	tvOn          bool
}

func NewTV(currentChan int, currentVolume int, tvOn bool) *SamsungTV {

	return &SamsungTV{currentChan: currentChan, currentVolume: currentVolume, tvOn: tvOn}
}

func (tv *SamsungTV) GetVolume() int {
	fmt.Println("SamsungTV volume is", tv.currentVolume)
	return tv.currentVolume
}
func (tv *SamsungTV) SetVolume(vol int) {
	fmt.Println("Setting SamsungTV volume to", vol)
	tv.currentVolume = vol
}

func (tv *SamsungTV) GetChannel() int {
	fmt.Println("SamsungTV channel is", tv.currentChan)
	return tv.currentChan
}
func (tv *SamsungTV) SetChannel(ch int) {
	fmt.Println("Setting SamsungTV channel to", ch)
	tv.currentChan = ch
}

func (tv *SamsungTV) SetOnState(tvOn bool) {
	if tvOn == true {
		fmt.Println("SamsungTV is on")
	} else {
		fmt.Println("SamsungTV is off")
	}
	tv.tvOn = tvOn
}
