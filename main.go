package main

import (
	"Anti-Brick/Hardware"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/olekukonko/tablewriter"
)

var Hmap = make(map[string]string)

const(
	a = "append"
	d = "delete"
	p = "print"
	P = "print_table"
)

func hashString(input string) string {
    hash := sha256.Sum256([]byte(input))
    return hex.EncodeToString(hash[:])
}

func Motherboard() string{
	command := Hardware.GetMotherBoardID()

	err, Output, _ := Hardware.RunCommand(command, Hardware.GetOS())

	if err != nil {
		fmt.Println(err)
	}else{
		Output := strings.Trim(Output, "\n")
		Output = strings.Replace(Output, "\r?\n", "", -1)
		return Output
	}
	return ""
}

func Ram() string{
	return Hardware.GetRamSerialNumber()
}

func Ssd() string{
	command := Hardware.GetSSDID()

	err, Output, _ := Hardware.RunCommand(command, Hardware.GetOS())

	if err != nil {
		fmt.Println(err)
	}else{
		Output := strings.Trim(Output, "\n")
		Output = strings.Replace(Output, "\r?\n", "", -1)
		return Output
	}
	return ""
}

func UUID() string{
	command := Hardware.GetUUID()

	err, Output, _ := Hardware.RunCommand(command, Hardware.GetOS())

	if err != nil {
		fmt.Println(err)
	}else{
		Output := strings.Trim(Output, "\n")
		return Output
	}
	return ""
}

func Battery() string{
	Hardware.GetBatterySerialNumber()
	html_file := "battery-report.html"
	serial_number := Hardware.GetBatteryNumber(html_file)
	serial_number = strings.Trim(serial_number, "\n")
	return serial_number
}

func HashMap(table *tablewriter.Table, key string, value string, flag string) {

	if flag == a{
		Hmap[key] = value
	}else if flag == d{
		delete(Hmap, key)
	}else if flag == P{
		for k,v := range Hmap{
			table.Append([]string{k, v, hashString(v)})
		}
		table.Render()
	}else if flag == p{
		for k,v := range Hmap{
			fmt.Printf("%s ---> %s", k, v)
		}
	}

}

func main(){	
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Component", "ID", "Hash"})

	start := time.Now()
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)
	ch5 := make(chan string)

	go func(){
		motherboard := Motherboard()
		ch1 <- motherboard
	}()
	go func(){
		ram := Ram()
		ch2 <- ram
	}()
	go func(){
		ssd := Ssd()
		ch3 <- ssd
	}()
	go func(){
		uuid := UUID()
		ch4 <- uuid
	}()
	go func(){
		battery := Battery()
		ch5 <- battery
	}()

	motherboard := <- ch1
	ram := <- ch2
	ssd := <- ch3
	uuid := <- ch4
	battery := <- ch5
	elapsed := time.Since(start)

	HashMap(table,"Motherboard", motherboard, a)
	HashMap(table ,"Ram", ram, a)
	HashMap(table ,"SSD", ssd, a)
	HashMap(table ,"UUID", uuid, a)
	HashMap(table ,"Battery", battery, a)
	fmt.Println(elapsed)

	var printstatus string
	fmt.Println("Enter 'p' to print the map")
	fmt.Scan(&printstatus)

	if printstatus == "p"{
		HashMap(table,"", "", p)
	}else if printstatus == "P"{
		HashMap(table,"", "", P)
	}

}