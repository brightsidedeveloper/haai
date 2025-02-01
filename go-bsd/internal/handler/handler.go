package handler

import (
	"server/internal/ai"
	"server/internal/bin"
	"server/internal/query"
	"server/internal/socket"
)

type Handler struct {
	bin   *bin.Bin
	query *query.Queries
	ss    *socket.Server
	ai    *ai.AI
}

func NewHandler(b *bin.Bin, q *query.Queries, ss *socket.Server, ai *ai.AI) *Handler {
	return &Handler{
		bin:   b,
		query: q,
		ss:    ss,
		ai:    ai,
	}
}
