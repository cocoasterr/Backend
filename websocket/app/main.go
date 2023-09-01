package main

import (
	"fmt"
	"net"
	"net/http"

	"golang.org/x/net/websocket"
)

type Message struct{
	Name string
	Price string
}

// untuk menampung connection di memory
type Hub struct{
	Clients map[*websocket.Conn]bool //untuk menyipan client dalam memory
	ClientRegisterChannel chan *websocket.Conn //ketika client masuk ke connection, connection clientnya disimpan melalui channel
	ClientRemovalChannel chan *websocket.Conn // menghapus ketika web socket ditutup, ini akan dihapus dari list of client, jadi tidak dapat broadcast messagenya
	BroadcastMessage chan Message //channel untuk broadcast message
}

//function ini akan dijalankan di routine yang berbeda
func (h *Hub) Run (){
	for{
		select{
		case conn := <- h.ClientRegisterChannel:
			h.Clients[conn] = true
		case conn := <- h.ClientRemovalChannel:
			delete(h.Clients, conn)
		case msg := <- h.BroadcastMessage:
			for conn:= range h.Clients{
				err := websocket.JSON.Send(conn, msg)
				if err != nil {
					fmt.Printf("error: %s\n", err)
				}
			}
		}
	}
}


func main(){
	h := &Hub{
		Clients: make(map[*websocket.Conn]bool),
		ClientRegisterChannel: make(chan *websocket.Conn),
		ClientRemovalChannel: make(chan *websocket.Conn),
		BroadcastMessage: make(chan Message),
	}
	go h.Run()

	http.Handle("/ws/bid", websocket.Handler(BidPrice(h)))


	host := GetHostName()
	fmt.Printf("Running on %s:8080", host)
	http.ListenAndServe(":8080", nil)
}

func BidPrice(h *Hub) func (*websocket.Conn){
	return func(conn *websocket.Conn){
		defer func(){
			h.ClientRemovalChannel <- conn
			conn.Close()
		}()
		name := conn.Request().URL.Query().Get("name")
		//lakukan register channel
		h.ClientRegisterChannel <- conn

		for{
			var receivedMessage Message
			// membaca setiap message yang datang
			err:= websocket.JSON.Receive(conn,&receivedMessage)			
			if err != nil {
				return 
			}
			msg:= Message{Name: name, Price: receivedMessage.Price}
			h.BroadcastMessage <- msg
		}
	}
}

func GetHostName() string{
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error:", err)
	}
	var host []string
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if ok && !ipAddr.IP.IsLoopback() && ipAddr.IP.To4() != nil {
			host = append(host, ipAddr.IP.String())
		}
	}
	return host[0]
}