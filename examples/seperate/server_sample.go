package main

import (
	"flag"
	"github.com/FreekingDean/gotWrap"
)

func main() {
	addr := flag.String("addr", "127.0.0.1", "Listening address")
	port := flag.String("port", "8000", "Listening port")
	pem := flag.String("pem", "certs/server.pem", "Cert pem file")
	key := flag.String("key", "certs/server.key", "Cert key file")
	flag.Parse()
	s := gotWrap.Server {
		Protocol: "tcp",
		ListenerAddr: *addr+":"+*port,
		PemFile: *pem,
		KeyFile: *key,
	}
	s.CreateServer()
}