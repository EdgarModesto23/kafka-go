package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		err := conn.SetDeadline(time.Now().Add(10 * time.Second))
		if err != nil {
			fmt.Println("Error setting deadline:", err)
			return
		}

		n, err := conn.Read(buf)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				fmt.Println("Read timeout, waiting for the next request...")
				return
			}
			fmt.Println("Error while reading from connection:", err.Error())
			return
		}
		req, errCode := ParseRequest(buf[:n])
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

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go handleConnection(conn)
	}
}
