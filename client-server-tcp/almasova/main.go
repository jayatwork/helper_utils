package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

var ipToScan = "10.12.58.143"
var minPort = 1
var maxPort = 1024
var timeOut = time.Duration(10) * time.Second

func main() {
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
}

func testTCPConnection(ip string, port int, doneChannel chan bool) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), timeOut)
	if err == nil {
		log.Printf("Host %s has open port %d", ip, port)
	}
	doneChannel <- true
}