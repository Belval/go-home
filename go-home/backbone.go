package main

import (
  "fmt"
  "net"
)

// Used to parse and save a response from
// a device following a Discover 'D' command
func parseAndSaveDiscoveryPacket([]byte content) {

}

// Used to parse and save a response from
// a device following a Get 'G' command.
func parseAndSaveGetPacket([]byte content) {

}

func sendUDPPacket([]byte content, string ipAddr, string port) (error) {
  conn, err = net.Dial("udp", ipAddr + ":" + port)
  if err != nil {
    return err
  }
  conn.Write(content)
  conn.Close()
  return nil
}

func parseUDPPacket([]byte content) {
  switch content[0] {
  case 'D':
  // This is a "Discovery" packet for new devices
    go parseAndSaveDiscoveryPacket(content)
  case 'G':
  // This is a "Get" response from a queried device
    go parseAndSaveGetPacket(content)
  default:
    return nil
  }
}

// Send a UDP "Discovery" packet to the
// broadcast address.
func sendDiscoverBroadcast() {
  sendUDPPacket("D:", "255.255.255.255", "30000")
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
