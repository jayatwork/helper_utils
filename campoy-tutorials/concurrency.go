package main

import (
 	"time"
 	 "net"
 	 "log"
 	 "io"
 	 "fmt"
 	 "strconv"
 )

 func triggerTalkAndSleep(talk string, d time.Duration) {
 	time.Sleep(d)
 	fmt.Println(talk)
 }

 func main() {
 	//Add code for autodetect on port
 	 var port int = 4002
 	 	l, err := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
 	 	if err != nil {
 	 		port = port + 1
 	 		log.Fatalf("Could not listen at HOST PORT given \n\n Trying next available port %v ", err)
 	 		fmt.Println()
 	 		m, err := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
 	 		if err != nil {
 	 			log.Fatalf("Could not listen at HOST PORT given \n\n FINAL ATTEMPT MADE %v ", err)
 	 		}

 	 		for {
 	 			connection2, err := m.Accept()
 	 			if err != nil {
 	 				log.Fatalf("Could not obtain a connection %v ", err)
 	 			}		
 	 			io.Copy(connection2, connection2)						
 	 			triggerTalkAndSleep("Ancrumblers", 3*time.Second)
 	 			triggerTalkAndSleep("there", 2*time.Second)
 	 			triggerTalkAndSleep("HIYA", 1*time.Second)
 	 		}

 	 	}
	
 	 for {
 	 	connection, err := l.Accept()
 	 	if err != nil {
 	 		log.Fatalf("Could not obtain a connection %v ", err)
 	 	}		
 	 	io.Copy(connection, connection)						
 		go triggerTalkAndSleep("Ancrumblers", 3*time.Second)
 		go triggerTalkAndSleep("there", 2*time.Second)
 		go triggerTalkAndSleep("HIYA", 1*time.Second)
 	}
}