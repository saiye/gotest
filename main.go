package main

import (
	"store/app/service/engine"
)

var server engine.FactoryServer
func main() {
	if len(server.ServerArr) == 0 {
		server = engine.FactoryServer{
			ServerArr: []engine.Server{
				engine.HttpServer{},
				engine.WebSocketServer{},
			},
		}
		server.Start()
	}
}
