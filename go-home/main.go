package main

import (
  "fmt"
)

func main() {
  fmt.Println("Loading configuration file...")
  loadConfig()
  fmt.Println("Insuring database state...")
  insureInitialDatabaseState()
	fmt.Println("Starting Backbone server...")
  startUDPServer(4567)
  fmt.Println("Starting Web server...")
  startWebServer(8080)
}
