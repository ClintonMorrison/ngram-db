package server

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"ngramdb/server/storage"
	"ngramdb/server/query"
	"ngramdb/server/database"
)


type Server struct {
	port string
	store *storage.Store
	database *database.Database
}

func New(port string) *Server {
	server := Server{}
	server.database = database.New()
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
	command := query.Parse(message)
	response := ""

	if command == nil {
		return "error\n"
	}

	response, err := s.database.AnswerQuery(command)
	if err != nil {
		return err.Error() + "\n"
	}

	return response + "\n"
}
