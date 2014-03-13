package main

import (
	"github.com/FreekingDean/gotWrap"
)

func main() {
	s := gotWrap.Server {
		Protocall: "tls",
		ListenerAddress: "127.0.0.0:8000",
		PemFile: "certs/server.pem",
		KeyFile: "certs/server.key",
	}
	go s.CreateServer()
	c := gotWrap.Client {
		Protocall: "tls",
		ListenerAddress: "127.0.0.0:8000",
		PemFile: "certs/client.pem",
		KeyFile: "certs/client.key",
	}
	c.Connect()
}
