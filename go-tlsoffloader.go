package main

import (
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net"
	"runtime"
)

var localAddress string
var backendAddress string
var certificatePath string
var keyPath string

func init() {
	flag.StringVar(&localAddress, "l", "localhost:44300", "local address")
	flag.StringVar(&backendAddress, "b", "gh-api.clickhouse.tech:9440", "backend address")
	// flag.StringVar(&certificatePath, "c", "cert.pem", "SSL certificate path")
	// flag.StringVar(&keyPath, "k", "key.pem", "SSL key path")
}

func main() {
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	listener, err := net.Listen("tcp", localAddress)
	if err != nil {
		log.Fatalf("error in tls.Listen: %s", err)
	}

	log.Printf("local server on: %s, backend server on: %s", localAddress, backendAddress)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error in listener.Accept: %s", err)
			break
		}
		log.Printf("local server on: %s, backend server on: %s", localAddress, backendAddress)

		log.Println("client " + conn.RemoteAddr().String() + " connected.")
		go handle(conn)
	}
}

func handle(clientConn net.Conn) {
	// rootca, err := loadRootCA("ca-certificates.crt")
	// if err != nil {
	// 	return nil, err
	// }

	config := tls.Config{
		// RootCAs: rootca,
		// InsecureSkipVerify: true
		// ServerName: "",
		// PreferServerCipherSuites
	}
	backendConn, err := tls.Dial("tcp", backendAddress, &config)

	if err != nil {
		log.Printf("error in tls.Dial: %s", err)
		clientConn.Close()
		return
	}

	go Tunnel(clientConn, backendConn)
	go Tunnel(backendConn, clientConn)
}

func Tunnel(from, to io.ReadWriteCloser) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recovered while tunneling")
		}
	}()

	io.Copy(from, to)
	to.Close()
	from.Close()
	log.Printf("tunneling is done")
}
