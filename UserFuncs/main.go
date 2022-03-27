package main

import (
	"fmt"
	"net"

	"github.com/KrunchyKernel/telnetServer/Server"
)

func main() {
	var connectedPorts = []int{8000}

	fmt.Println(connectedPorts[0])
	var listeners []net.Listener

	fmt.Println(connectedPorts, listeners)
	err := Server.StartServer(&connectedPorts, &listeners)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}
	fmt.Println(listeners)
}
