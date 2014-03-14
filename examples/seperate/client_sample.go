package main

import (
	"flag"
	"strconv"
	"github.com/FreekingDean/gotWrap"
)

func main() {
	addr := flag.String("addr", "127.0.0.1", "Remote address")
	port := flag.String("port", "8000", "Remote port")
	pem := flag.String("pem", "certs/client.pem", "Cert pem file")
	key := flag.String("key", "certs/client.key", "Cert key file")
	flag.Parse()
	c := gotWrap.Client {
		Protocol: "tcp",
		RemoteAddr: *addr+":"+*port,
		PemFile: *pem,
		KeyFile: *key,
	}
	c.Connect()
	for i:=0;i<100;i++ {

		c.SendMessage("Hello"+strconv.Itoa(i)+"\n")
	}
}