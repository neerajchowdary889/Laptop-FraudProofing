package Hardware


func GetSSDID() string {
	os := GetOS()
	switch os{
	case "windows":
		command := "wmic diskdrive get serialnumber"
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
