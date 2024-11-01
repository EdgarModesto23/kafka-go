package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	buf := make([]byte, 1024)

	// Read in a non-blocking way
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Connection timeout, closing...")
			return // Exit the goroutine if the context is done
		default:
			// Attempt to read from the connection
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Error while reading from connection: ", err.Error())
				return
			}

			// Process the request
			req, errCode := ParseRequest(buf[:n]) // Use only the bytes read
			if errCode != 0 {
				res := Response{corr_id: req.headers.corr_id, length: 0, err_code: int16(errCode)}
				conn.Write(ResponseToByte(res))
				return
			}

			response := GetHandler(*req).Execute()
			if _, err := conn.Write(ResponseToByte(response)); err != nil {
				fmt.Print(err)
				return
			}
		}
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}
  
	for {
		// Accept a new connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue // Handle error and continue accepting more connections
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}
