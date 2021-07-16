package port

/*
	Really ropey port scanner in Go ( POC!! )
	1 - This just fires up loads of noisy threads, an IDS will see this coming a mile off!
	2 - The routine handling is crap, just waiting off the end for it to be done. needs some WGs
	3 - Only does tcp right now and only does up to ports 1024...very slowly!
	4 - My first attempt at something useful using channels and routines!
*/

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type TargetProtoHost struct {
	Proto string
	Host  string
	Port  int
}

func ScanPort(protocol, hostname string, port int) {

	// result := ScanResult{Port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 10*time.Second)
	if err == nil {
		defer conn.Close()
		fmt.Printf(" %v : %v : %v : %v\n", protocol, hostname, port, "OPEN")
	}

}

func PortThread(pid int, jobs <-chan TargetProtoHost) {
	// we need to keep feeding this "thread" by constantly sucking data off the pipe
	// if you don't the thread stalls and you get channel deadlock
	for tmpVar := range jobs {
		ScanPort(tmpVar.Proto, tmpVar.Host, tmpVar.Port)
	}
}

func InitialScan(hostname string) {

	jobThreads := 50

	jobs := make(chan TargetProtoHost, jobThreads)

	for t := 1; t <= jobThreads; t++ {
		go PortThread(t, jobs)
	}

	for i := 1; i <= 1024; i++ {
		jobs <- TargetProtoHost{
			Proto: "tcp",
			Host:  hostname,
			Port:  i,
		}
	}
	close(jobs)

	time.Sleep(5 * time.Second)

}
