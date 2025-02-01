package routes

import (
	"server/internal/handler"

	"github.com/go-chi/chi/v5"
)

func MountRoutes(r chi.Router, h *handler.Handler) {

	ai := chi.NewRouter()
	ai.Post("/prompt", h.PostPrompt)

	game := chi.NewRouter()
	game.Get("/rooms", h.GetRooms)

	api := chi.NewRouter()
	api.Mount("/ai", ai)
	api.Mount("/game", game)

	api.Get("/users", h.GetUsers)
	api.Post("/user", h.PostUser)

	r.Mount("/api", api)
	r.Get("/ws", h.HandleConnections)
}
