package Hardware

import (
	"github.com/StackExchange/wmi"
)

func GetRamSerialNumber() string {
	var modules []struct {
		SerialNumber string
	}
	query := "SELECT SerialNumber FROM Win32_PhysicalMemory"
	err := wmi.Query(query, &modules)
	if err != nil {
		return ""
	}
	if len(modules) > 0 {
		return modules[0].SerialNumber
	}
	return ""
}
