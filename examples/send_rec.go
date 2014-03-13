package main

import (
	"github.com/FreekingDean/gotWrap"
)

func main() {
	s := gotWrap.Server {
		protocall: "tls",
		listenerAddress: "127.0.0.0:8000",
		pemFile: "certs/server.pem",
		keyFile: "certs/server.key",
	}
	go s.CreateServer()
	c := gotWrap.Client {
		protocall: "tls",
		listenerAddress: "127.0.0.0:8000",
		pemFile: "certs/client.pem",
		keyFile: "certs/client.key",
	}
	c.Connect()
}
