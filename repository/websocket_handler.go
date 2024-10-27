package repository

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "sync"
)

type WebSocketHub struct {
    clients   map[*websocket.Conn]bool
    broadcast chan []byte
    mu        sync.Mutex
}

var Hub = WebSocketHub{
    clients:   make(map[*websocket.Conn]bool),
    broadcast: make(chan []byte),
}

func (h *WebSocketHub) HandleConnections(c *gin.Context) {
    conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 0, 0)
    if err != nil {
        return
    }
    defer conn.Close()

    h.mu.Lock()
    h.clients[conn] = true
    h.mu.Unlock()

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            h.mu.Lock()
            delete(h.clients, conn)
            h.mu.Unlock()
            break
        }
        h.broadcast <- msg
    }
}

func (h *WebSocketHub) Broadcast(message []byte) {
    h.mu.Lock()
    defer h.mu.Unlock()
    for client := range h.clients {
        err := client.WriteMessage(websocket.TextMessage, message)
        if err != nil {
            client.Close()
            delete(h.clients, client)
        }
    }
}