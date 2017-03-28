	package main

	import (
		"fmt"
		"log"
		"net"
		"encoding/gob"
			)

	type P struct{N,M string}

	func main() {
		fmt.Println("start client at localhost:5005");
		conn, err := net.Dial("tcp", "localhost:5005")
		if err != nil 	{ 
			log.Fatal("Connection error", err)
						}
		encoder := gob.NewEncoder(conn)
		p := P{"1","2"}
		encoder.Encode(p)
		fmt.Println("M =",p.M);
		fmt.Println("N =",p.N);
		conn.Close()
		fmt.Println("Values sent to server");
				} 