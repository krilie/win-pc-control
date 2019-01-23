package win_control

import (
	"errors"
	"github.com/go-vgo/robotgo"
)

//play stop pause vol_up vol_down mute next prev POST
func Media(action string) error {
	switch action {
	case "play":
		robotgo.KeyTap("audio_play")
	case "pause":
		robotgo.KeyTap("audio_pause")
	case "stop":
		robotgo.KeyTap("audio_stop")
	case "next":
		robotgo.KeyTap("audio_next")
	case "prev":
		robotgo.KeyTap("audio_prev")
	case "vol_up":
		robotgo.KeyTap("audio_vol_up")
	case "vol_down":
		robotgo.KeyTap("audio_vol_down")
	case "mute":
		robotgo.KeyTap("audio_mute")
	default:
		return errors.New("参数错误:" + action)
	}
	return nil
}
