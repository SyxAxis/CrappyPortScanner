package main

import (
	"fmt"

	"github.com/syxaxis/portScan/port"
)

func main() {

	fmt.Println("PortScan")
	port.InitialScan("192.168.0.10")

}
