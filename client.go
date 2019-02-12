package nchats

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	server *server
	conn   *websocket.Conn
	send   chan message
}

func (c *client) readPump() {
	defer func() {
		c.server.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		var msg message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		msg.Time = time.Now().Format("2006-01-02 15:04:05")

		//c.server.hub.broadcast <- msg

		var msgBytes bytes.Buffer
		err = json.NewEncoder(&msgBytes).Encode(msg)
		if err != nil {
			log.Printf("error: %v", err)
		}
		c.server.nconn.Publish("nchats.message", msgBytes.Bytes())
	}
}

func (c *client) writePump() {
	defer func() {
		c.conn.Close()
	}()

	for {
		msg, ok := <-c.send
		if !ok {
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		c.conn.WriteJSON(msg)
	}
}

func (s *server) wsHandler() http.HandlerFunc {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		c := &client{
			server: s,
			conn:   conn,
			send:   make(chan message),
		}
		s.hub.register <- c

		// start go rutines for sending and receiving messages from websocket
		go c.writePump()
		go c.readPump()
	}
}
