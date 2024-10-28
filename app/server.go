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

  _, headers, _, err_code := ParseRequest(buf)
  if err_code != 0 {
    conn.Write(Response(0, headers.corr_id, int32(err_code)))
    return
  }
  
  conn.Write(Response(0, headers.corr_id, 0))

}
