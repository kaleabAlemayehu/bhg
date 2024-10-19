package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	dest, err := net.Dial("tcp", ":8081")
	if err != nil {
		log.Fatalln(err.Error())
	}
	go func() {
		_, err = io.Copy(dest, conn)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()
	for {
		_, err = io.Copy(conn, dest)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}
