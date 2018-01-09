package main

import (
  "fmt"
)

func main() {
	fmt.Println("Starting Backbone server...")
  startUDPServer(4567)
  fmt.Println("Starting Web server...")
  startWebServer(8080)
}
