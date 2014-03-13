package main

import (
	"time"
	"github.com/FreekingDean/gotWrap"
)

func main() {
	s := gotWrap.Server {
		Protocall: "tcp",
		ListenerAddress: "127.0.0.1:8000",
		PemFile: "certs/server.pem",
		KeyFile: "certs/server.key",
	}
	go s.CreateServer()
	time.Sleep(1000*time.Millisecond)
	c := gotWrap.Client {
		Protocall: "tcp",
		RemoteAddr: "127.0.0.1:8000",
		PemFile: "certs/client.pem",
		KeyFile: "certs/client.key",
	}
	c.Connect()
}
