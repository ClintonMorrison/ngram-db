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
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Port     int
	database *database.Database
	handler  *handler.QueryHandler
	filename string
}

func New(port int, filename string) *Server {
	server := Server{}

	db, err := database.FromFile(filename)
	if err != nil {
		fmt.Printf("Could not read from database file %s. Got error %s\n", filename, err.Error())
		db = database.New()
	}

	server.database = db
	server.handler = handler.New(server.database)
	server.Port = port
	server.filename = filename

	return &server
}

func (s *Server) Listen() {
	s.saveOnExit()
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

func (s *Server) saveOnExit() {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v\n", sig)
		fmt.Println("Saving database to file")
		s.database.ToFile(s.filename)
		time.Sleep(2*time.Second)
		os.Exit(0)
	}()
}
