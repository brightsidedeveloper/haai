package handler

import (
	"net/http"
	"server/internal/buf"
)

func (h *Handler) PostPrompt(w http.ResponseWriter, r *http.Request) {
	var body buf.AIRequest
	if err := h.bin.UnmarshalBody(r.Body, &body); err != nil {
		h.bin.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	resp, err := h.ai.Prompt(body.Prompt)
	if err != nil {
		h.bin.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.bin.ProtoWrite(w, http.StatusOK, &buf.AIResponse{Response: resp})
}
