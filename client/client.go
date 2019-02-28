package client

import (
	"net"
	"bufio"
	"fmt"
	"strings"
)

type Client struct {
	address string
	logging bool
	conn net.Conn
}


func New(address string, logging bool) *Client {
	client := Client{}
	client.address = address
	client.logging = logging

	return &client
}

func (c *Client) log(action string, msg string) {
	if !c.logging {
		return
	}

	fmt.Printf("> %s: %s\n", action, msg)
}

func (c *Client) Connect() error {
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		return err
	}

	c.log("Connected", c.address)

	c.conn = conn
	return nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Send(message string) (string, error) {
	c.log("Send", message)

	_, err := c.conn.Write([]byte(message + "\n"))
	if err != nil {
		return "", err
	}

	response, err := bufio.NewReader(c.conn).ReadString('\n')
	c.log("Receive", response)

	return strings.TrimSpace(response), err
}
