package main

import (
	"fmt"
	"net"
)

func main() {
	var listener net.Listener                             // Declare a variable to hold the listener
	var err error                                         // Declare a variable to hold any error that occurs
	listener, err = net.Listen("tcp","127.0.0.1:9000") //assign the listener to listen on TCP port
	//check for the error and terminate program after printing the error if present
	if err != nil {
		fmt.Println("error detected", err)
		return
	}
	defer listener.Close() //release the listener when the function exits
	fmt.Println("server is bound to port and listening")
}
