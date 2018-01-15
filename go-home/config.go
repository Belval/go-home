package main

import (
  "os"
  "encoding/json"
)

var (
  config *Config
)

type Config struct {
  EncrytionKey string `json:"ENCRYPTION_KEY"`
  Database struct {
    Driver string `json:"DRIVER"`
    User string `json:"USER"`
    Password string `json:"PASSWORD"`
    Dbname string `json:"DBNAME"`
    Host string `json:"HOST"`
    Port string `json:"PORT"`
  } `json:"DATABASE_CONFIG"`
}

func loadConfig() (error) {
  var config Config
  confFile, err := os.Open("config.json")
  if err != nil {
    return err
  }
  jsonParser := json.NewDecoder(confFile)
  jsonParser.Decode(&config)
  return nil
}
