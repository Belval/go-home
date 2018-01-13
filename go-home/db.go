package main

import (
  _ "database/sql"
  "fmt"
  _ "github.com/mattn/go-sqlite3"
  "github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE home_devices (
    id INT,
    uuid VARCHAR(64),
    tag VARCHAR(64),
    last_heartbeat DATETIME,
    last_ip VARCHAR(64),
);

CREATE TABLE home_data (
    id INT,
    device_id INT,
    temperature INT,
    light INT,
    movement BIT,
    is_on BIT,
    timestamp DATETIME,
)

CREATE TABLE home_users (
    id INT,
    firstname VARCHAR(64),
    lastname VARCHAR(64),
    telephone VARCHAR(32),
    email VARCHAR(64),
    password_hash VARCHAR(256),
    password_salt VARCHAR(256),
)
`

type HomeDevice struct {
  Id int `db:"id"`
  Uuid string `db:"uuid"`
  Tag string `db:tag`
  LastHeartbeat string `db:last_heartbeat`
  LastIp string `db:last_ip`
}

type HomeData struct {
  Id int `db:"id"`
  DeviceId int `db:"device_id"`
  Temperature int `db:"temperature"`
  Light int `db:"light"`
  Movement bool `db:"movement"`
  IsOn bool `db:"is_on"`
  Timestamp string `db:"timestamp"`
}

type HomeUsers struct {
  Id int `db:"id"`
  Firstname string `db:"firstname"`
  Lastname string `db:"lastname"`
  Telephone string `db:"telephone"`
  Email string `db:"email"`
  PasswordHash string `db:"password_hash"`
  PasswordSalt string `db:"password_salt"`
}

func insureInitialDatabaseState() (error) {
  // We test the connection
  db, err := getDatabaseConnection()
  if err != nil {
    return err
  }
  // We execute our schema
  db.MustExec(schema)
  return nil
}

func getDatabaseConnection() (*sqlx.DB, error) {
  driver, user, password, dbname := "a", "b", "c", "d"
  // We test the connection
  db, err := sqlx.Connect(driver, fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname))
  if err != nil {
    return nil, err
  }
  return db, nil
}

func listHomeDevices() (error) {
  return nil
}

func addHomeDevice(homeDevice HomeDevice) (error) {
  db, err := getDatabaseConnection()
  if err != nil {
    return err
  }
  tr := db.MustBegin()
  tr.NamedExec("INSERT INTO  (uuid, tag, last_heartbeat, last_ip) VALUES (:uuid, :tag, :last_heartbeat, :last_ip)", homeDevice)
  tr.Commit()
  return nil
}

func deleteHomeDevice(id int) (error) {
  return nil
}

func registerUser() (error) {
  return nil
}

func disableUser() (error) {
  return nil
}
