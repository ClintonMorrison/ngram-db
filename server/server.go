package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"ngramdb/server/database"
	"ngramdb/server/handler"
	"ngramdb/server/query"
	"strings"
)

type Server struct {
	Port     int
	database *database.Database
	handler  *handler.QueryHandler
}

func New(port int) *Server {
	server := Server{}
	server.database = database.New()
	server.handler = handler.New(server.database)
	server.Port = port

	return &server
}

func (s *Server) Listen() {
	listener, err := net.Listen("tcp4", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go s.handleRequest(connection)
	}
}

func (s *Server) handleRequest(connection net.Conn) {
	fmt.Printf("Serving %s\n", connection.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(connection).ReadString(';')
		if err != nil {
			fmt.Println(err)
			return
		}

		message := strings.TrimSpace(string(netData))
		message = strings.TrimRight(message, ";")
		if message == "STOP" {
			break
		}

		response := s.processMessage(message)
		connection.Write([]byte(string(response + "\n")))
	}
	connection.Close()
}

func (s *Server) processMessage(message string) string {
	query, err := query.Parse(message)
	response := s.handler.Handle(query, err)
	return asJSON(response)
}

func asJSON(obj interface{}) string {
	serialized, err := json.Marshal(obj)

	if err != nil {
		return "{\"success\": false}"
	}

	return string(serialized)
}
