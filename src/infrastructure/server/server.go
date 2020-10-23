package server

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"

	"leagueapi.com.br/brain/src/entities"

	"leagueapi.com.br/brain/src/infrastructure/syncronize"

	"github.com/gorilla/websocket"
)

// Server brains server
type Server struct {
	Addr   *string
	URL    *url.URL
	Conn   *websocket.Conn
	Sync   *syncronize.Syncronize
	Header http.Header
	Brain  *entities.Brain
}

func (s *Server) dial() {

	c, _, err := websocket.DefaultDialer.Dial(s.URL.String(), s.Header)
	if err != nil {
		log.Fatal("dial:", err)
	}

	s.Conn = c
}

// Handler handle all socket talk
func (s *Server) Handler() {
	s.Brain.AddSyncronize(s.Sync)
	for {
		select {
		case <-s.Sync.Done:
			return
		case commmand := <-s.Sync.NewMessage:
			s.Brain.Handle(commmand)
		case observerUpdate := <-s.Sync.ObserverUpdate:
			err := s.Conn.WriteMessage(websocket.TextMessage, []byte(observerUpdate))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-s.Sync.Interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := s.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			return
		}
	}

}

// Reciver is a reciver new msg
func (s *Server) Reciver() {
	defer close(s.Sync.Done)
	for {
		_, message, err := s.Conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
		s.Sync.NewMessage <- string(message)
	}
}

// BindHeader bind header
func (s *Server) BindHeader() *Server {
	s.Header["Ws-Client-Name"] = []string{"brain"}
	return s
}

// BindURL create url to connect
func (s *Server) BindURL() *Server {
	s.URL = &url.URL{Scheme: "ws", Host: *s.Addr, Path: "/live"}
	return s
}

// Connect dialer ws
func (s *Server) Connect() *Server {
	s.dial()
	return s
}

// NewServer server constructor
func NewServer() *Server {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost:9000"
	}

	return &Server{
		Addr:   flag.String("addr", host, "http service address"),
		Sync:   syncronize.NewSyncronize(),
		Header: http.Header{},
		Brain:  entities.NewBrain().StartHandlers(),
	}
}
