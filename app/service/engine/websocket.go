package engine

import (
	"fmt"
)

type WebSocketServer struct {
}

func (server WebSocketServer) Start() {

}
func (server WebSocketServer) Exit() {
	fmt.Println("websocket Exit")
}
