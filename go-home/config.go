package main

import (

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

func loadConfig() (*Config, error) {
  return nil, nil
}
