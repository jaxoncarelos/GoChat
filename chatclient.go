package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func filter[K comparable, V any](data map[K]V, f func(K) bool) map[K]V {
	fltd := make(map[K]V)
	for i, e := range data {
		if f(i) {
			fltd[i] = e
		}
	}

	return fltd
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Message Received: ", string(p))
		for i, ws := range filter(listOfConnects, func(s string) bool { return s != conn.RemoteAddr().String() }) {
			fmt.Println("Sending to ", i)
			if err := ws.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

var listOfConnects = make(map[string]*websocket.Conn)

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("New connection from ", r.RemoteAddr)
	listOfConnects[r.RemoteAddr] = ws
	reader(ws)
}
func setupRoutes() {
	http.HandleFunc("/ws", wsEndpoint)
}
func main() {
	fmt.Println("Hello, World!")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
