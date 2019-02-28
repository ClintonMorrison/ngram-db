package main

import (
	"ngramdb/server"
	// "ngramdb/client"
	// "fmt"
)

func main() {
	// host := "localhost"
	port := "3000"
	// address := fmt.Sprintf("%s:%s", host, port)

	s := server.New(port)
	s.Listen()

	/*
	c := client.New(address, true)
	c.Connect()
	c.Send("ADD SET test(3);")
	c.Send("ADD TEXT(\"ABCD\") TO test;")
	c.Send("ADD TEXT(\"AB\") TO test;")
	c.Send("GET COUNT OF \"AB\" IN test;")
	*/
}
