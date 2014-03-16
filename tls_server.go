package gotWrap

import (
	"crypto/tls"
	"log"
)

type callBack func(*tls.Conn, string)

type Server struct {
	ListenerAddr string
	Protocol string
	PemFile string
	KeyFile string
	MessageRec callBack
}

func (server *Server) CreateServer() {
	//TODO - Auto gen certs upon first start
	cert, err := tls.LoadX509KeyPair(server.PemFile, server.KeyFile)
	if err != nil {
		log.Fatalf("[gotWrap-SERVER] loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}, ClientAuth: tls.RequireAnyClientCert}
	
	listener, err := tls.Listen(server.Protocol, server.ListenerAddr, &config)
	if err != nil {
		log.Fatalf("[gotWrap-SERVER] listening on: %s :%s", listener.Addr().String(), err)
	}
	log.Print("[gotWrap-SERVER] listening")
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("[gotWrap-SERVER] accept: %s", err)
			break
		}
		log.Printf("[gotWrap-SERVER] accepted from %s", conn.RemoteAddr())
		tlscon, ok := conn.(*tls.Conn)
		if ok && handshake(tlscon) {
			go handleClient(tlscon, server.MessageRec)
		} else {
			conn.Close()
			log.Printf("[gotWrap-SERVER] conn: closed")
		}
	}
}

func handleClient(tlscon *tls.Conn, mcb callBack) {
	defer tlscon.Close()
	log.Print("[gotWrap-SERVER] conn: type assert to TLS succeedded")
	buf := make([]byte, 512)
	for {
		log.Print("[gotWrap-SERVER] conn: waiting")
		n, err := tlscon.Read(buf)
		if err != nil {
			log.Printf("[gotWrap-SERVER] conn: read err: %s", err)
			break
 		}
 		log.Printf("[gotWrap-SERVER] conn: read: %q", string(buf[:n]))
 		mcb(tlscon, string(buf[:n]))		
	}
	log.Println("[gotWrap-SERVER] server: conn: closed")
}

func handshake(tlscon *tls.Conn) bool {
	err := tlscon.Handshake()
	if err != nil {
		log.Fatalf("[gotWrap-SERVER] handshake failed: %s", err)
		return false
	} else {
		log.Print("[gotWrap-SERVER] conn: Handshake completed")
	}
	state := tlscon.ConnectionState()
	log.Println("[gotWrap-SERVER] mutual: ", state.NegotiatedProtocolIsMutual)
	return true
}

func SendMessage(tlscon *tls.Conn, buf[] byte) {
	log.Printf("[gotWrap-SERVER] conn: write: %q", string(buf))
	n, err := tlscon.Write(buf)
	log.Printf("[gotWrap-SERVER] conn: wrote %d bytes", n)
	if err != nil {
		log.Printf("[gotWrap-SERVER] write-failed:%s", err)
		tlscon.Close()
	}
}