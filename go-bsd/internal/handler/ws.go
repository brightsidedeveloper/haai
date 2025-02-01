package handler

import (
	"fmt"
	"net/http"
	"server/internal/socket"
)

func (h *Handler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := socket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	c := socket.NewClient(conn, h.ss)

	go c.ReadMessages()
}
