package gotWrap

import (
    "crypto/tls"
    "io"
    "log"
)

type Client struct {
    RemoteAddr string
    Protocol string
    PemFile string
    KeyFile string
    conn *tls.Conn
}

func (client *Client) Connect() {
    cert, err := tls.LoadX509KeyPair(client.PemFile, client.KeyFile)
    if err != nil {
        log.Fatalf("server: loadkeys: %s", err)
    }
    config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
    client.conn, err = tls.Dial(client.Protocol, client.RemoteAddr, &config)
    if err != nil {
        log.Fatalf("client: dial: %s", err)
    }
    log.Println("client: connected to: ", client.conn.RemoteAddr())
    state := client.conn.ConnectionState()
    log.Println("client: handshake: ", state.HandshakeComplete)
    log.Println("client: mutual: ", state.NegotiatedProtocolIsMutual)
    
    go client.listen()
}

func (client *Client) SendMessage(m string) {
    message := "Hello\n"
    n, err := io.WriteString(client.conn, message)
    if err != nil {
        log.Fatalf("client: write: %s", err)
        client.conn.Close()
    }
    log.Printf("client: wrote %s (%d bytes)", message, n)
}

func (client *Client) listen() {
    defer client.conn.Close()
    for {
        reply := make([]byte, 256)
        n, err := client.conn.Read(reply)
        if err != nil {
            log.Fatalf("client: dial: %s", err)
        }
        log.Printf("client: read %s (%d bytes)", string(reply[:n]), n)
    }
}