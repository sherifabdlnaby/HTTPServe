package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	defer listener.Close()

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Accepted a connection...")

		go ServeConn(conn)
	}
}

func ServeConn(conn net.Conn) {

	//Close Connection after Serve
	defer conn.Close()

	for {
		// Refresh Deadline
		conn.SetDeadline(time.Now().Add(time.Second * 30))

		// Create Buffer and Read
		var buf [128]byte
		n, err := conn.Read(buf[:])

		if err != nil {
			log.Printf("Finished Connection")
			return
		}

		os.Stderr.Write(buf[:n])
	}
}
