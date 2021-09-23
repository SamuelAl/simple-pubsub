package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	handleError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	handleError(err)
	fmt.Printf("Started listening on port %s\n", service)
	for {
		handleMessage(conn)
	}
}

func handleMessage(conn *net.UDPConn) {
	// Create buffer.
	var buff [512]byte
	n, addr, err := conn.ReadFromUDP(buff[0:])
	if err != nil {
		return
	}
	temp := readTemp(n, buff)
	fmt.Printf("Received reading from %v: temp = %q\n", addr, temp)
	conn.WriteToUDP([]byte("ACK"), addr)
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

func readTemp(n int, buff [512]byte) string {
	msg := string(buff[0:n])
	return msg[strings.Index(msg, "/")+1:]
}
