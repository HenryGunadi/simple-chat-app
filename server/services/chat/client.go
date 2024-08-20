package chat

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	id   string
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	maxMessageSize = 512
	pingPeriod = (pongWait * 9) / 10
)

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading the connection : %v", err)
		return
	}
	log.Println("Ws connection is established")

	id := uuid.New().String()

	client := &Client{
		id: id,
		hub: hub,
		conn: ws,
		send: make(chan []byte),
	}

	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func(){
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			_, err = w.Write(msg);
			if err != nil {
				log.Println("error write messages : ", err)
				return
			}
			msgStr := string(msg)
			log.Println("message successfully sent : ", msgStr)


			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(msg)	
			}

			if err := w.Close(); err != nil {
				return
			}
		
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) readPump() {
	defer func() {
		c.conn.Close()
		c.hub.unregister <- c
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(appData string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, text, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error read message: %v", err)
			}
			break
		}

		log.Printf("Received raw message: %s", text)

		msg := &WsMessage{}
		err = json.Unmarshal(text, msg)
		if err != nil {
			log.Println("error decoding text messages : ", err)
			continue
		}

		log.Println("Decoded messages : ", msg.Text)
		c.hub.broadcast <- &Message{ClientID: c.id, Text: msg.Text}
	}
}







