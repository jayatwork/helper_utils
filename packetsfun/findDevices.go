package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	//Find all devices

	var devices []pcap.Interface
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	//Print all device information
	fmt.Println("Devices Found.\n")

	for _, device := range devices {
		fmt.Println("Device Name: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println("Device Addresses: ")
		for _, address := range device.Addresses {
			fmt.Println("IP Address: ", address.IP)
			fmt.Println("IP Address: ", address.Netmask)
		}
	}
}
