package main

import (
  "fmt"
  "net"
)

func sentUDPPacket([]byte content, string ipAddr, string port) (error) {
  fmt.Println(string(content))
  return nil
}

func parseUDPPacket([]byte content) (error) {
  fmt.Println(string(content))
  return nil
}

func startUDPServer(port int) {
  udpConn, err := net.ListenUDP("udp", ":" + string(port))
	if err != nil {
		fmt.Println(err)
	}

	for {
    buffer := [2048]byte{}
    offset, err := udpConn.Read(buffer[0:])
    if err != nil {
      fmt.Println("Unable to read buffer!")
    }
    parseUDPPacket(buffer[0:offset])
	}
}
