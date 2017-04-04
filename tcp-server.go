package main

import (
    "fmt"
    "net"
    "encoding/json"
)
type Message struct {  
    Value string `json:"Value"`
}

func handleConnection(conn net.Conn) {
    dec := json.NewDecoder(conn)
    a := &Message{}
    dec.Decode(a)
	fmt.Println("Values received from client");
    fmt.Println("Value =",a.Value);
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