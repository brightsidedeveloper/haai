package socket

import (
	"fmt"
	"log"
	"server/internal/buf"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	conn   *websocket.Conn
	roomId string
	server *Server
	id     string
	name   string
}

func NewClient(conn *websocket.Conn, s *Server) *Client {
	return &Client{
		server: s,
		conn:   conn,
		id:     uuid.New().String(),
	}
}

func (c *Client) ReadMessages() {
	defer func() {
		c.server.mut.Lock()
		delete(c.server.Clients, c.id)
		c.server.mut.Unlock()
	}()

	fmt.Println("Client connected")

	for {
		msgType, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		if msgType != websocket.BinaryMessage {
			continue
		}

		var wsMsg buf.WSMessage
		if err := proto.Unmarshal(msg, &wsMsg); err != nil {
			log.Println(err)
			continue
		}

		switch wsMsg.Type {
		case buf.SocketMessageType_SET_NAME:
			c.name = wsMsg.GetSetName().GetName()
			continue
		case buf.SocketMessageType_CREATE_ROOM:
			name := wsMsg.GetCreateRoom().GetName()

			exists := false
			for _, r := range c.server.Rooms {
				if r.Name == name {
					fmt.Println("Room already exists")
					c.SendProtoMessage(&buf.WSMessage{
						Type: *buf.SocketMessageType_ERROR.Enum(),
						Payload: &buf.WSMessage_Error{
							Error: &buf.Error{
								Message: "Room already exists",
							},
						},
					})
					exists = true
					break
				}
			}
			if exists {
				continue
			}

			fmt.Println("Creating room")
			room := NewRoom(name)
			c.server.mut.Lock()
			c.server.Rooms[room.Id] = room
			c.server.mut.Unlock()
			c.roomId = room.Id
			room.AddClient(c)
			continue
		case buf.SocketMessageType_JOIN_ROOM:
			roomId := wsMsg.GetJoinRoom().GetId()
			room, ok := c.server.Rooms[roomId]
			if !ok {
				continue
			}
			c.roomId = roomId
			room.AddClient(c)
			fmt.Println("Client joined room")
			fmt.Println("Room clients:", len(room.Clients))
			continue
		case buf.SocketMessageType_LEAVE_ROOM:
			continue
		}

	}
}

func (c *Client) SendProtoMessage(message proto.Message) {
	data, err := proto.Marshal(message)
	if err != nil {
		log.Println("Failed to serialize Protobuf:", err)
		return
	}
	c.conn.WriteMessage(websocket.BinaryMessage, data)
}
