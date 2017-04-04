package main

import (  
    "encoding/json"
    "fmt"
	"log"
    "net"
)

type Person struct {  
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
}

func main() {  
	fmt.Println("start client at localhost:5005");
		conn, err := net.Dial("tcp", "localhost:5005")
		if err != nil 	{ 
			log.Fatal("Connection error", err)
						}
		encoder := json.NewEncoder(conn)
    in := `{"firstName":"John","lastName":"Dow"}`

    rawIn := json.RawMessage(in)
    bytes, err := rawIn.MarshalJSON()
    if err != nil {
        panic(err)
    }

    var p Person
    err = json.Unmarshal(bytes, &p)
    if err != nil {
        panic(err)
    }
	encoder.Encode(p)
	conn.Close()
}