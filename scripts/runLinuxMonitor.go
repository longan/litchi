package main

import (
	"fmt"
	"github.com/longan/litchi/linuxmonitor"
)

func main() {
	fmt.Println("Start linux monitor")

	stat, _ := linuxmonitor.GetCpu()

	fmt.Println(stat.CPUStatAll)

}