package engine

import (
	"os"
	"os/signal"
)

type FactoryServer struct {
	ServerArr []Server
}

func (fac *FactoryServer) Start() {
	for _, v := range fac.ServerArr {
		go v.Start()
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	for _, server := range fac.ServerArr {
		 server.Exit()
	}
}
