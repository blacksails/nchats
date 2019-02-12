package nchats

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nats-io/go-nats"
)

type server struct {
	r     *chi.Mux
	hub   *hub
	nconn *nats.Conn
}

type Options struct {
	Port    int
	NATSURL string
}

func Start(opts Options) error {

	hub := &hub{
		clients:    make(map[*client]bool),
		register:   make(chan *client),
		unregister: make(chan *client),
		broadcast:  make(chan message),
	}

	nconn, err := nats.Connect(opts.NATSURL)
	if err != nil {
		return err
	}

	s := &server{
		r:     chi.NewMux(),
		hub:   hub,
		nconn: nconn,
	}

	s.nconn.Subscribe("nchats.message", func(natsMsg *nats.Msg) {
		var msg message
		json.NewDecoder(bytes.NewBuffer(natsMsg.Data)).Decode(&msg)
		s.hub.broadcast <- msg
	})

	s.r.Use(middleware.Logger)
	s.r.Get("/ws", s.wsHandler())
	s.r.Get("/*", s.staticHandler())

	// start client hub
	go s.hub.run()

	return http.ListenAndServe(fmt.Sprintf(":%v", opts.Port), s.r)
}

func (s *server) staticHandler() http.HandlerFunc {
	fs := http.FileServer(http.Dir("app/dist"))
	return func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}
}
