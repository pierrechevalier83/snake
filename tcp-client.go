package main

import (  
    "encoding/json"
    "fmt"
	"log"
    "net"
)

type Message struct {  
    Value string `json:"Value"`
}

func main() {  
	fmt.Println("start client at localhost:5005");
		conn, err := net.Dial("tcp", "localhost:5005")
		if err != nil 	{ 
			log.Fatal("Connection error", err)
						}
		encoder := json.NewEncoder(conn)
    in := `{"Value":"42"}`

    rawIn := json.RawMessage(in)
    bytes, err := rawIn.MarshalJSON()
    if err != nil {
        panic(err)
    }

    var p Message
    err = json.Unmarshal(bytes, &p)
    if err != nil {
        panic(err)
    }
	encoder.Encode(p)
	conn.Close()
}