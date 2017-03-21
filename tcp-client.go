package main

import (
    "fmt"
    "log"
    "net"
    "encoding/gob"
)

type P struct {
   N int64
}

func main() {
    fmt.Println("start client");
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal("Connection error", err)
    }
    encoder := gob.NewEncoder(conn)
    p := P{42}
    encoder.Encode(p)
	fmt.Println("sent ",p.N);
    conn.Close()
    fmt.Println("done");
}