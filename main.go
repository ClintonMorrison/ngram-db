package main

import (
	"fmt"
	"ngramdb/client"
	"ngramdb/server"
	"flag"
	"bufio"
	"os"
)

func main() {
	isServer := flag.Bool(
		"server",
		false,
		"Run the ngramdb server")

	isClient := flag.Bool(
		"client",
		false,
		"Run the ngramdb client")

	port := flag.Int(
		"port",
		3000,
		"Port of server (defaults to 3000)")

	host := flag.String(
		"host",
		"localhost",
		"Hostname of server (defaults to localhost)")

	flag.Parse()

	if !*isServer && !*isClient {
		flag.Usage()
	}

	if *isServer && *isClient {
		go runServer(*port)
	} else if *isServer {
		runServer(*port)
	}

	if *isClient {
		runClient(*host, *port)
	}
}

func runServer(port int) {
	s := server.New(port)
	s.Listen()
}

func runClient(host string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)
	c := client.New(address, true)
	c.Connect()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		query, _ := reader.ReadString('\n')
		response, err := c.Send(query)
		fmt.Println(response, err)
	}
}