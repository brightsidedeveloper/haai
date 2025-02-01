package socket

import "sync"

type Server struct {
	Clients map[string]*Client
	Rooms   map[string]*Room
	mut     sync.Mutex
}

func NewServer() *Server {
	return &Server{
		Clients: make(map[string]*Client),
		Rooms:   make(map[string]*Room),
	}
}
