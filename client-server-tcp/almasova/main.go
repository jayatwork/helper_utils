package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

var ipToScan = "192.168.1.1"
var minPort = 1
var maxPort = 1024
var timeOut = time.Duration(10) * time.Second

func main() {

	// Stage one PORT_SCANNER functionality
	activeThreads := 0
	doneChannel := make(chan bool)

	for port := minPort; port <= maxPort; port++ {
		go testTCPConnection(ipToScan, port, doneChannel)
		activeThreads++
	}

	for activeThreads > 0 {
		<- doneChannel
		activeThreads--
	}

	//Stage two IDENTIFY_SERVICES_AVAIALABLE

}

func testTCPConnection(ip string, port int, doneChannel chan bool) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), timeOut)
	if err == nil {
		log.Printf("Host %s has open port %d", ip, port)
	}
	doneChannel <- true
}