package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	handleError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	handleError(err)

	for {
		temp := getTemperature()
		_, err = conn.Write([]byte(fmt.Sprintf("temp/%f", temp)))
		handleError(err)

		var buff [512]byte
		n, err := conn.Read(buff[0:])
		handleError(err)
		fmt.Println(string(buff[0:n]))
		time.Sleep(2 * time.Second)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %v", err.Error())
		os.Exit(1)
	}
}

func getTemperature() float32 {
	temp := (rand.Float32() * 10) + 30
	return temp
}
