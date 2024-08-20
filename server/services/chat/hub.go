package chat

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	ClientID string
	Text     string
}

type WsMessage struct {
	Text    string      `json:"text"`
	Headers map[string]string `json:"HEADERS"`
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
	messages 	[]*Message
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true

			log.Println("client registered : ", client.id)
			for key, value := range h.clients {
				fmt.Printf("clientID : %v, value: %v", key, value)
			}
			for _, msg := range h.messages {
				client.send <- getMessage(msg)
			}
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				log.Println("client unregistred : ", client.id)
				close(client.send)
				delete(h.clients, client)
			}
		case msg := <-h.broadcast:
			h.messages = append(h.messages, msg)

			for client := range h.clients {
				select {
				case client.send <- getMessage(msg):
					
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func getMessage(msg *Message) []byte {
	// Convert Message to JSON
    message := WsMessage{
        Text: msg.Text,
        Headers: map[string]string{
            "ClientID": msg.ClientID,
        },
    }
    jsonData, err := json.Marshal(message)
    if err != nil {
        log.Println("Error marshalling message:", err)
        return nil
    }
	
    return jsonData
}