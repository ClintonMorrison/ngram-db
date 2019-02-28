package tests

/*
import (
	"fmt"
	"testing"
	"go-redis-clone/client"
	"go-redis-clone/server"
)


type testCase func (c *client.Client) func(*testing.T)

func expectResponse(
	t *testing.T,
	c *client.Client,
	message string,
	expectedResponse string) {
	response, err := c.Send(message)
	if err != nil {
		t.Error(err)
	}

	if response != expectedResponse {
		t.Error(fmt.Sprintf("Expected '%s' but got '%s'", expectedResponse, response))
	}
}

func TestSystem(t *testing.T) {
	host := "localhost"
	port := "3000"
	address := fmt.Sprintf("%s:%s", host, port)

	s := server.New(port)
	go s.Listen()

	c := client.New(address, true)
	c.Connect()

	testCases := map[string]testCase{
		"FLUSHALL": testFlushAll,
		"PING": testPing,
		"SET and GET": testSetGet,
		"KEYS": testKeys,
		"RENAME": testRename,
	}

	for name, testCase := range testCases {
		c.Send("FLUSHALL")
		t.Run(name, testCase(c))
	}
}

func testPing(c *client.Client) func(*testing.T) {
	return func(t *testing.T) {
		expectResponse(t, c, "PING", "PONG")
	}
}

func testSetGet(c *client.Client) func(*testing.T) {
	return func(t *testing.T) {
		expectResponse(t, c, "SET test abcd", "")
		expectResponse(t, c, "GET test", "abcd")
	}
}

func testKeys(c *client.Client) func(*testing.T) {
	return func(t *testing.T) {
		c.Send("SET test1 abcd")
		c.Send("SET test2 abcd")
		expectResponse(t, c, "KEYS", "[\"test1\",\"test2\"]")
	}
}

func testFlushAll(c *client.Client) func(*testing.T) {
	return func(t *testing.T) {
		c.Send("SET test1 abcd")
		expectResponse(t, c, "FLUSHALL", "")
		expectResponse(t, c, "KEYS", "[]")
	}
}

func testRename(c *client.Client) func(*testing.T) {
	return func(t *testing.T) {
		c.Send("SET test1 abcd")
		expectResponse(t, c, "RENAME test1 test2", "")
		expectResponse(t, c, "KEYS", "[\"test2\"]")
		expectResponse(t, c, "GET test2", "abcd")
	}
}

*/

