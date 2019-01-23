package win_control

import (
	"errors"
	"fmt"
	"os/exec"
)

var errParam = errors.New("没有找到对应动作。。。")

func winShutdown() error {
	cmd := exec.Command("shutdown", "/s", "/t", "1")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// on off
func winMonitor(action string) error {
	switch action {
	case "on":
		cmd := exec.Command("./win10bin/monitor-control-on.exe")
		if err := cmd.Run(); err != nil {
			return err
		}
	case "off":
		cmd := exec.Command("./win10bin/monitor-control-off.exe")
		if err := cmd.Run(); err != nil {
			return err
		}
	default:
		fmt.Println("sys_control", "no action on monitor")
		return errParam
	}
	return nil
}

func SysCtl(action string) error {
	switch action {
	case "shutdown":
		return winShutdown()
	case "monitor_on":
		return winMonitor("on")
	case "monitor_of":
		return winMonitor("off")
	default:
		return errParam
	}
}
