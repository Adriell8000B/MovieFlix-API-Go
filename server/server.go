package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(server *http.Server) *Server {
	return &Server{
		server: server,
	}
}

func (s *Server) setup() {}

func (s *Server) listen() {
	listen_err := s.server.ListenAndServe()
	if listen_err != nil {
		fmt.Println("Couldn't start server, and error ocurred: ", listen_err)
	}

	fmt.Println("Look mom, it's alive")
}

func (s *Server) Init() {
	s.setup()
	s.listen()
}
