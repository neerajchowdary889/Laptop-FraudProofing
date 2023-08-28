package main

import (
	"Anti-Brick/Hardware"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/jedib0t/go-pretty/table"
)

var Hmap = make(map[string]string)

const(
	a = "append"
	d = "delete"
	p = "print"
	t = "print_table"
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

func HashMap(Table *table.Table, key string, value string, flag string) {

	if flag == a{
		Hmap[key] = value
	}else if flag == d{
		delete(Hmap, key)
	}else if flag == t{
		for k,v := range Hmap{
			row := table.Row{k, v, hashString(v)}
			Table.AppendRow(row)
		}
		Table.Render()
	}else if flag == p{
		for k,v := range Hmap{
			fmt.Printf("%s ---> %s", k, v)
		}
	}

}

func main(){	

	Table := table.NewWriter()
	Table.SetOutputMirror(os.Stdout)
	Table.AppendHeader(table.Row{"Component", "ID", "Hash"}) 

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

	HashMap(Table.(*table.Table),"Motherboard", motherboard, a)
	HashMap(Table.(*table.Table) ,"Ram", ram, a)
	HashMap(Table.(*table.Table) ,"SSD", ssd, a)
	HashMap(Table.(*table.Table) ,"UUID", uuid, a)
	HashMap(Table.(*table.Table) ,"Battery", battery, a)
	fmt.Println(elapsed)

	var printstatus string
	fmt.Println(">>>Enter 'p' to print the map\n>>>Enter 't' to print the table")
	fmt.Scan(&printstatus)
	printstatus = strings.ToLower(strings.TrimSpace(printstatus))

	if printstatus == "p"{
		HashMap(Table.(*table.Table),"", "", p)
	}else if printstatus == "t"{
		HashMap(Table.(*table.Table),"", "", t)
	}
}