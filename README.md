this first code reserves port 9000 and allows the local host to access it only.


package main

import (
	"fmt"
	"net"
)

func main() {
	var listener net.Listener                             // Declare a variable to hold the listener
	var err error                                         // Declare a variable to hold any error that occurs
	listener, err = net.Listen("tcp", "127.0.0.1:9000") //assign the listener to listen on TCP port
	//check for the error and terminate program after printing the error if present
	if err != nil {
		fmt.Println("error detected", err)
		return
	}
	defer listener.Close() //release the listener when the function exits
	fmt.Println("server is bound to port and listening")
}

this code inside main function accepts connection from a single client

// accepting single client connection
	var conn net.Conn             // connection holding variable
	conn, err = listener.Accept() // wait for connection to be established
	//check for the error and terminate program after printing the error if present
	if err != nil {
		fmt.Println("error detected", err)
		return
	}
	defer conn.Close() //release the connection after function exits
	fmt.Println("connection accepted from", conn.RemoteAddr())
	

	INSIDE MAIN:

	package main

import (
	"fmt"
	"net"
)

func main() {
	var listener net.Listener                           // Declare a variable to hold the listener
	var err error                                       // Declare a variable to hold any error that occurs
	listener, err = net.Listen("tcp", "127.0.0.1:9000") //assign the listener to listen on TCP port
	//check for the error and terminate program after printing the error if present
	if err != nil {
		fmt.Println("error detected", err)
		return
	}
	defer listener.Close() //release the listener when the function exits
	fmt.Println("server is bound to port and listening")

    // accepting single client connection
	var conn net.Conn             // connection holding variable
	conn, err = listener.Accept() // wait for connection to be established
	//check for the error and terminate program after printing the error if present
	if err != nil {
		fmt.Println("error detected", err)
		return
	}
	defer conn.Close() //release the connection after function exits
	fmt.Println("connection accepted from", conn.RemoteAddr())
}