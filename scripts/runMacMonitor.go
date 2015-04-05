package main

import (
	"fmt"
	"github.com/longan/litchi/macmonitor"
)

func main() {
	fmt.Println("Start to monitor mac")

	metricsString := macmonitor.GetMetrics()

	fmt.Println(metricsString)
}
