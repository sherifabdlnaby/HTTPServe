package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

	// Deadline TIMEOUT
	conn.SetDeadline(time.Now().Add(time.Second * 30))

	// Create Buffer and Read
	scanner := bufio.NewScanner(conn)

	// Parse Header
	headers := parseHeader(scanner)

	// Respond
	responseBody := "HELLO WORLD <br/> You requested: " + headers["uri"]
	writeResponse(conn, responseBody)

	return
}

func writeResponse(conn net.Conn, responseBody string) {
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	io.WriteString(conn, "content-type: text/html; charset=UTF-8")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(responseBody))
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, responseBody)
}

func parseHeader(scanner *bufio.Scanner) map[string]string {
	// parse
	headers := make(map[string]string)
	// Scan Request-Line
	if scanner.Scan() {
		headers["request-line"] = scanner.Text()
		requestLineFields := strings.Fields(scanner.Text())
		headers["method"] = requestLineFields[0]
		headers["uri"] = requestLineFields[1]
		headers["http-version"] = requestLineFields[2]
	}
	for scanner.Scan() {
		line := scanner.Text()

		keyIndex := strings.IndexRune(line, ':')
		if keyIndex == -1 {
			break //TODO Invalid Headers error
		}

		headers[line[:keyIndex]] = line[keyIndex+2:]
	}
	return headers
}
