package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)


type Server struct{
	conn map[*websocket.Conn]bool
}
func NewServer()*Server{
	return &Server{
		conn: make(map[*websocket.Conn]bool),
	}
}

func (s *Server)HandleWS(ws *websocket.Conn){
	fmt.Println("new incoming connection from clents: ", ws.RemoteAddr())
	s.conn[ws] = true
	s.readLoop(ws)
}

func (s *Server)readLoop(ws *websocket.Conn){
	buf := make([]byte,1024)
	for{
		n,err := ws.Read(buf)
		if err == io.EOF{
			break
		}
		if err != nil {
			fmt.Println("err :", err)
			continue
		}
		msg := buf[:n]
		fmt.Println(string(msg))
		ws.Write([]byte("Thanks for the message!"))
	}

}

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.HandleWS))
	http.ListenAndServe(":3000", nil)
}