package Hardware

import (
	"runtime"
	"github.com/abdfnx/gosh"
)

func GetOS() string {
	return runtime.GOOS
}

func RunCommand(command string, os string) (error, string, string){

	switch os{
		case "windows":
			err,out, _ := gosh.PowershellOutput(command)
			return err, out, ""
		case "linux":
			err,out, _ := gosh.ShellOutput(command)
			return err, out, ""
		case "darwin":
			err,out, _ := gosh.ShellOutput(command)
			return err, out, ""
		default:
			return nil, "false", "false"	
	}
}

func GetMotherBoardID() string {
	os := GetOS()
	switch os{
	case "windows":
		command := "wmic bios get serialnumber"
		return command
	case "linux":
		command := "hal-get-property --udi /org/freedesktop/Hal/devices/computer --key system.hardware.uuid"
		return command
	case "darwin":
		command := "ioreg -l | grep IOPlatformSerialNumber"
		return command
	default:
		return "false"
	}
	return "false"
}
