package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"ngramdb/responses"
	"strings"
)

type Client struct {
	address string
	logging bool
	conn    net.Conn
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

	_, err := c.conn.Write([]byte(message + ";"))
	if err != nil {
		return "", err
	}

	response, err := bufio.NewReader(c.conn).ReadString('\n')
	c.log("Receive", response)

	return strings.TrimSpace(response), err
}

func (c *Client) GetSets() (*responses.Sets, error) {
	query := fmt.Sprintf("GET SETS")
	ptr, err := c.sendAndParse(query, &responses.Sets{})
	response, ok := ptr.(*responses.Sets)

	if !ok {
		return nil, err
	}

	return response, err
}


func (c *Client) AddSet(setName string, n int) (*responses.Generic, error) {
	query := fmt.Sprintf("ADD SET %s(%d)", setName, n)
	ptr, err := c.sendAndParse(query, &responses.Generic{})
	response, ok := ptr.(*responses.Generic)

	if !ok {
		return nil, err
	}

	return response, err
}

func (c *Client) DeleteSet(setName string) (*responses.Generic, error) {
	query := fmt.Sprintf("DELETE SET %s", setName)
	ptr, err := c.sendAndParse(query, &responses.Generic{})
	response, ok := ptr.(*responses.Generic)

	if !ok {
		return nil, err
	}

	return response, err
}

func (c *Client) AddText(setName string, text string) (*responses.Generic, error) {
	query := fmt.Sprintf("ADD TEXT '%s' IN %s", text, setName)
	ptr, err := c.sendAndParse(query, &responses.Generic{})
	response, ok := ptr.(*responses.Generic)

	if !ok {
		return nil, err
	}

	return response, err
}

func (c *Client) GetNGrams(setName string, n int) (*responses.NGrams, error) {
	query := fmt.Sprintf("GET NGRAMS(%d) IN %s", setName, n)
	ptr, err := c.sendAndParse(query, &responses.NGrams{})
	response, ok := ptr.(*responses.NGrams)

	if !ok {
		return nil, err
	}

	return response, err
}

func (c *Client) GetCount(setName string, ngram string) (*responses.Count, error) {
	query := fmt.Sprintf("GET COUNT OF '%s' IN %s", ngram, setName)
	ptr, err := c.sendAndParse(query, &responses.Count{})
	response, ok := ptr.(*responses.Count)

	if !ok {
		return nil, err
	}

	return response, err
}

func (c *Client) GetFrequency(setName string, ngram string) (*responses.Frequency, error) {
	query := fmt.Sprintf("GET FREQ OF '%s' IN %s", ngram, setName)
	ptr, err := c.sendAndParse(query, &responses.Frequency{})
	response, ok := ptr.(*responses.Frequency)

	if !ok {
		return nil, err
	}

	return response, err
}

func (c *Client) GetCompletions(setName string, pattern string) (*responses.Completions, error) {
	query := fmt.Sprintf("GET COMPLETIONS OF '%s' IN %s", pattern, setName)
	ptr, err := c.sendAndParse(query, &responses.Completions{})
	response, ok := ptr.(*responses.Completions)

	if !ok {
		return nil, err
	}

	return response, err
}

func (c *Client) GetProbableSet(ngram string) (*responses.ProbableSet, error) {
	query := fmt.Sprintf("GET PROBABLE SET OF '%s'", ngram)
	ptr, err := c.sendAndParse(query, &responses.ProbableSet{})
	response, ok := ptr.(*responses.ProbableSet)

	if !ok {
		return nil, err
	}

	return response, err
}

func (c *Client) sendAndParse(query string, response interface{}) (interface{}, error) {
	rawResponse, err := c.Send(query)
	if err != nil {
		return nil, err
	}

	err = parseResponse(rawResponse, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func parseResponse(rawResponse string, responseObj interface{}) error {
	bytes := []byte(rawResponse)
	err := json.Unmarshal(bytes, responseObj)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, responseObj)
	if err != nil {
		return err
	}

	// Return error if response contained error
	errorResponse := &responses.Error{}
	err = json.Unmarshal(bytes, errorResponse)
	if errorResponse.ErrorType != "" {
		return errorResponse
	}

	return nil
}