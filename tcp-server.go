package main

import (
    "fmt"
    "net"
    "encoding/json"
)
type P struct {
     N,M string
}

func handleConnection(conn net.Conn) {
    dec := json.NewDecoder(conn)
    a := &P{}
    dec.Decode(a)
	fmt.Println("Values received from client");
    fmt.Println("M =",a.M);
	fmt.Println("N =",a.N);
    conn.Close()
}

func main() {
    fmt.Println("start server at localhost:5005 ");
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