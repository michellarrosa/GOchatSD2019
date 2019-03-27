package main

import (
    "fmt",
    "net",
    "log"
)

func main(conn) {
  fmt.Fprintf(conn, "Hello Caller!")
  // do more stuff with conn
  conn.Close()
}
