package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	buf := make([]byte, 1024)

	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Error while reading from connection: ", err.Error())
		os.Exit(1)
	}

	req, err_code := ParseRequest(buf)
	if err_code != 0 {
		res := Response{corr_id: req.headers.corr_id, length: 0, err_code: int16(err_code)}
		conn.Write(ResponseToByte(res))
		return
	}

	response := GetHandler(*req).Execute()

  _, err = conn.Write(ResponseToByte(response))
  if err != nil {
    fmt.Print(err)
    os.Exit(1)
  }
}
