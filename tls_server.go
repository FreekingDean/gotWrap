package gotWrap

import (
	"net"
	"crypto/tls"
	"log"
)

type callBack func(net.Addr, string)

type Server struct {
	ListenerAddr string
	Protocol string
	PemFile string
	KeyFile string
	MessageRec callBack
	connections map[net.Addr]*tls.Conn //map[RemoteAddress]TLS_Connection
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
		if ok && server.handshake(tlscon.RemoteAddr()) {
			server.connections[tlscon.RemoteAddr()] = tlscon
			go server.handleClient(tlscon.RemoteAddr())
		} else {
			conn.Close()
			log.Printf("[gotWrap-SERVER] conn: closed")
		}
	}
}

func (server *Server) handleClient(conn net.Addr) {
	defer server.connections[conn].Close()
	log.Print("[gotWrap-SERVER] conn: type assert to TLS succeedded")
	buf := make([]byte, 512)
	for {
		log.Print("[gotWrap-SERVER] conn: waiting")
		n, err := server.connections[conn].Read(buf)
		if err != nil {
			log.Printf("[gotWrap-SERVER] conn: read err: %s", err)
			break
 		}
 		log.Printf("[gotWrap-SERVER] conn: read: %s", string(buf[:n]))
 		server.MessageRec(conn, string(buf[:n]))		
	}
	delete(server.connections, conn)
	log.Println("[gotWrap-SERVER] server: conn: closed")
}

func (server *Server) handshake(conn net.Addr) bool {
	err := server.connections[conn].Handshake()
	if err != nil {
		log.Fatalf("[gotWrap-SERVER] handshake failed: %s", err)
		return false
	} else {
		log.Print("[gotWrap-SERVER] conn: Handshake completed")
	}
	state := server.connections[conn].ConnectionState()
	log.Println("[gotWrap-SERVER] mutual: ", state.NegotiatedProtocolIsMutual)
	return true
}

func (server *Server) SendMessage(conn net.Addr, buf[] byte) {
	log.Printf("[gotWrap-SERVER] conn: write: %s\n", string(buf))
	n, err := server.connections[conn].Write(buf)
	log.Printf("[gotWrap-SERVER] conn: wrote %d bytes", n)
	if err != nil {
		log.Printf("[gotWrap-SERVER] write: %s", err)
		server.connections[conn].Close()
	}
}