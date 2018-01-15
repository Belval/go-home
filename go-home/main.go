package main

import (
  "fmt"
)

func main() {
  fmt.Println("Loading configuration file...")
  err := loadConfig()
  if err != nil {
    fmt.Println("Unable to load config.json (does it exist?)")
  }
  fmt.Println("Insuring database state...")
  insureInitialDatabaseState()
	fmt.Println("Starting Backbone server...")
  go startUDPServer(4567)
  fmt.Println("Starting Web server...")
  startWebServer(8080)
}
