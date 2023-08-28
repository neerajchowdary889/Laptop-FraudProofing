package Hardware

func GetUUID() string {
	os := GetOS()
	switch os{
	case "windows":
		command := "wmic csproduct get uuid"
		return command
	default:
		return "false"
	}
	return "false"
}
