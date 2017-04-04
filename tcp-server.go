package main

import (
    "fmt"
    "net"
    "encoding/json"
)
type Person struct {  
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
}

func handleConnection(conn net.Conn) {
    dec := json.NewDecoder(conn)
    a := &Person{}
    dec.Decode(a)
	fmt.Println("Values received from client");
    fmt.Println("N =",a.FirstName);
    conn.Close()
}

func main() {
    fmt.Println("start");
   ln, err := net.Listen("tcp", ":5005")
    if err != nil {
        // handle error
    }
    for {
        conn, err := ln.Accept() // this blocks until connection or error
        if err != nil {
            // handle error
            continue
        }
        go handleConnection(conn) // a goroutine handles conn so that the loop can accept other connections
    }
}