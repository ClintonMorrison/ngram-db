package tests


import (
	"fmt"
	"testing"
	"ngramdb/client"
	"ngramdb/server"
)


type testCase func (c *client.Client) func(*testing.T)


func TestSystem(t *testing.T) {
	host := "localhost"
	port := 3000


	testCases := map[string]testCase{
		"SETS": testBasicSetOperations,
		"NGRAMS": testBasicNGramQueries,
		"ADVANCED": testAdvancedNGramQueries,
	}

	// Spin up a server for each test case
	servers := make(map[string]*server.Server, len(testCases))
	for name, _ := range testCases {
		servers[name] = server.New(port, "")
		go servers[name].Listen()
		port++
	}


	for name, testCase := range testCases {
		server := servers[name]
		address := fmt.Sprintf("%s:%d", host, server.Port)

		c := client.New(address, true)
		err := c.Connect()

		if err != nil {
			panic(err)
		}

		fmt.Printf("RUNNING TEST CASE %s ON %s", name, address)
		t.Run(name, testCase(c))
	}
}

func testBasicSetOperations(c *client.Client) func(*testing.T) {
	return func(t *testing.T) {
		// Add a set
		_, err := c.AddSet("test", 3)
		assertNoError(t, err)

		// Expect duplicate key error
		_, err = c.AddSet("test", 3)
		assertError(t, err)

		// Get sets
		setsResponse, err := c.GetSets()
		if setsResponse.Sets[0] != "test" {
			t.Error("Expected set to be added")
		}

		// Remove a set
		_, err = c.DeleteSet("test")
		assertNoError(t, err)

		// Expect no such set error
		_, err = c.DeleteSet("test")
		assertError(t, err)


		// Expect set removed
		setsResponse, err = c.GetSets()
		assertNoError(t, err)
		if len(setsResponse.Sets) != 0 {
			t.Error("Expected set to be deleted")
		}
	}
}

func testBasicNGramQueries(c *client.Client) func(*testing.T) {
	return func(t *testing.T) {
		// Add a set
		_, err := c.AddSet("test", 3)
		assertNoError(t, err)

		// Add text
		_, err = c.AddText("test", "AABC")
		assertNoError(t, err)

		// Expect no such set error
		_, err = c.AddText("test1", "AABC")
		assertError(t, err)

		// Get count
		countResponse, err := c.GetCount("test", "A")
		assertNoError(t, err)
		if countResponse.Count != 2 {
			t.Errorf("Expected count to be 2 but got %d", countResponse.Count)
		}

		// Get frequency
		frequencyResponse, err := c.GetFrequency("test", "A")
		assertNoError(t, err)
		if frequencyResponse.Frequency != 0.5 {
			t.Error("Expected frequency to be 0.5")
		}
	}
}

func testAdvancedNGramQueries(c *client.Client) func(*testing.T) {
	return func(t *testing.T) {
		c.AddSet("set_a", 1)
		c.AddText("set_a", "AAAA")

		c.AddSet("set_b", 1)
		c.AddText("set_b", "BBBB")

		response, err := c.GetProbableSet("AA")
		assertNoError(t, err)
		if response.Set != "set_a" {
			t.Error("Expected set_a to be most similar")
		}
	}
}


func assertNoError(t *testing.T, e error) {
	if e != nil {
		t.Error("Unexpected error: " + e.Error())
	}
}

func assertError(t *testing.T, e error) {
	if e == nil {
		t.Error("Expected error")
	}
}
