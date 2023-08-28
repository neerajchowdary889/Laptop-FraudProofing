package Hardware

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func GetBatterySerialNumber() string {

    command := "powercfg /batteryreport"

    err, output, _ := RunCommand(command, GetOS())

    if err != nil {
        return ""
    }
    for _, line := range strings.Split(string(output), "\n") {
        if strings.Contains(line, "Serial Number") {
            return strings.TrimSpace(strings.Split(line, ":")[1])
        }
    }
    return ""
}

func GetBatteryNumber(htmlFile string) string {
    htmlContent, err := ioutil.ReadFile(htmlFile)
    if err != nil {
        return err.Error()
    }
    serialNumberRegex:= regexp.MustCompile(`<td><span class="label">SERIAL NUMBER</span></td><td>(.*?)</td>`)
	
    match := serialNumberRegex.FindStringSubmatch(string(htmlContent))
    if len(match) > 1 {
        return match[1]
    } else {
        return ""
    }
}

