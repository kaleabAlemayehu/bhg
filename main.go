package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port!")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		con, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go echo(con)
	}
}

func echo(conn net.Conn) {
	defer conn.Close()
	b := make([]byte, 512)
	for {
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected Error")
			break
		}
		log.Printf("Received %d bytes: %s", size, string(b))
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to Write data")
		}
	}
}
