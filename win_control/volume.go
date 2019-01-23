package win_control

import (
	"os/exec"
	"strconv"
)

// -2 -1 0-100
func WinVolume(value int) error {
	cmd := exec.Command("./win10bin/volume.exe", strconv.Itoa(value))
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
