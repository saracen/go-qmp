package qmp

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Message struct {
	Return json.RawMessage `json:"return"`
	Error  *Error          `json:"error"`
}

type Error struct {
	Class string `json:"class"`
	Desc  string `json:"desc"`
}

func (e Error) Error() string {
	return fmt.Sprintf("class:%s, desc: %s", e.Class, e.Desc)
}

type messageHandler struct {
	mu      sync.Mutex
	pending map[int64]chan Message
}

func newMessageHandler() *messageHandler {
	return &messageHandler{
		pending: make(map[int64]chan Message),
	}
}

func (h *messageHandler) register(id int64) chan Message {
	h.mu.Lock()
	defer h.mu.Unlock()

	ch := make(chan Message, 1)
	h.pending[id] = ch

	return ch
}

func (h *messageHandler) handle(resp response) bool {
	h.mu.Lock()
	handler, ok := h.pending[resp.ID]
	if ok {
		delete(h.pending, resp.ID)
	}
	h.mu.Unlock()

	if ok {
		handler <- resp.Message
		close(handler)
		return true
	}

	return false
}

func (h *messageHandler) clear() {
	h.mu.Lock()
	defer h.mu.Unlock()

	for id, handler := range h.pending {
		close(handler)
		delete(h.pending, id)
	}
}
