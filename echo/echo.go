package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err.Error())
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = io.Copy(conn, conn)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

}
