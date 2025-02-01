package socket

import (
	"sync"

	"github.com/google/uuid"
)

type Room struct {
	Id        string
	Name      string
	Clients   map[string]*Client
	Game      *Game
	Broadcast chan []byte
	mut       sync.Mutex
}

func NewRoom(name string) *Room {
	return &Room{
		Id:      uuid.New().String(),
		Name:    name,
		Clients: make(map[string]*Client),
		Game:    newGame(),
	}
}

func (r *Room) AddClient(c *Client) {
	r.mut.Lock()
	r.Clients[c.id] = c
	r.mut.Unlock()
}

func (r *Room) RemoveClient(c *Client) {
	c.roomId = ""
	r.mut.Lock()
	delete(r.Clients, c.id)
	r.mut.Unlock()
}
