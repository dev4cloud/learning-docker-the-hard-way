package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

const port = 2000
const protocol = "tcp"

func main() {
	server, err := net.Listen(protocol, ":"+strconv.Itoa(port))
	if server == nil {
		fmt.Fprintf(os.Stderr, "Error starting Copycat: %v\n", err.Error())
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Copycat listening at %v ...\n", server.Addr())
	connectionChannel := makeConnectionChannel(server)
	for {
		go handleConnection(<-connectionChannel)
	}
}

func makeConnectionChannel(listener net.Listener) chan net.Conn {
	channel := make(chan net.Conn)
	go func() {
		for {
			client, err := listener.Accept()
			if client == nil {
				fmt.Fprintf(os.Stderr, "Error accepting incoming connection: %v\n", err.Error())
				continue
			}
			fmt.Fprintf(os.Stdout, "Accepted incoming connection from: %v\n", client.RemoteAddr())
			channel <- client
		}
	}()
	return channel
}

func handleConnection(client net.Conn) {
	io.Copy(client, client)
	client.Close()
}
