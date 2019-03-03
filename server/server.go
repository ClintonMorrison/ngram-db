package server

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"ngramdb/server/query"
	"ngramdb/server/database"
	"ngramdb/server/handler"
	"encoding/json"
)


type Server struct {
	port string
	database *database.Database
	handler *handler.QueryHandler
}

func New(port string) *Server {
	server := Server{}
	server.database = database.New()
	server.handler = handler.New(server.database)
	server.port = port

	return &server
}

func (s *Server) Listen() {
	listener, err := net.Listen("tcp4", ":" + s.port)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
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
		if message == "STOP" {
			break
		}

		response := s.processMessage(message)
		connection.Write([]byte(string(response + "\n")))
	}
	connection.Close()
}

func (s *Server) processMessage(message string) string {
	fmt.Println("Recieved " + message)
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