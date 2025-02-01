package handler

import (
	"net/http"
	"server/internal/buf"
)

func (h *Handler) GetRooms(w http.ResponseWriter, r *http.Request) {

	rooms := make([]*buf.Room, 0)
	for id, room := range h.ss.Rooms {
		rooms = append(rooms, &buf.Room{
			Id:   id,
			Name: room.Name,
		})
	}

	h.bin.ProtoWrite(w, http.StatusOK, &buf.GetRooms{
		Rooms: rooms,
	})
}
