	package main

	import (
		"fmt"
		"log"
		"net"
		"encoding/json"
			)

	type Message struct {
		Value string
	}

	func main() {
		fmt.Println("start client at localhost:5005");
		conn, err := net.Dial("tcp", "localhost:5005")
		if err != nil 	{ 
			log.Fatal("Connection error", err)
						}
		encoder := json.NewEncoder(conn)
		const jsonStream = `{"Value": "42"}`
		encoder.Encode(jsonStream)
		fmt.Println("jsonStream=",jsonStream);
		conn.Close()
		fmt.Println("Values sent to server");
				} 
				