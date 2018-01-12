package home_db

import (
  "database/sql"
  "fmt"
  "time"
  "github.com/mattn/go-sqlite3"
)

/*
  go-home DB structure is simple,
  home_devices: Device informations
	=> id (INT)
        => uuid (VARCHAR(64))
        => tag (VARCHAR(64))
        => last_heartbeat (DATETIME)
        => last_ip (VARCHAR(64))
  home_data: Data returned by the Get command
        => id (INT)
	=> device_id (INT)
        => temperature (INT)
        => light (INT)
        => movement (BIT)
        => is_on (BIT)
        => timestamp ()
  home_users: User table
        => firstname (VARCHAR(64))
        => lastname (VARCHAR(64))
        => telephone (VARCHAR(32))
        => email (VARCHAR(64))
        => password_hash (VARCHAR(256))
        => password_salt (VARCHAR(256))
*/

func checkIfDbIsBuilt() (bool, error) {
  return false, nil
}

func listHomeDevices() () {
}

func addHomeDevice() () {
}

func deleteHomeDevice() () {
}


