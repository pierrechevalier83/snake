package main

import (
    "fmt"
    "net"
    "encoding/gob"
)
type P struct {
     N,M string
}

func handleConnection(conn net.Conn) {
    dec := gob.NewDecoder(conn)
    a := &P{}
    dec.Decode(a)
	fmt.Println("Values received from client");
    fmt.Println("M =",a.M);
	fmt.Println("N =",a.N);
    conn.Close()
}

func main() {
    fmt.Println("start");
   ln, err := net.Listen("tcp", ":8080")
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