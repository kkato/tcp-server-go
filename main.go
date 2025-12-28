package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {

	for {
		time.Sleep(time.Second * 3)
		buf := make([]byte, 100)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		requestInfo := string(buf[:n])

		fmt.Print("request: ")
		fmt.Println(requestInfo)

		if requestInfo == "\"close\"" {
			fmt.Println("closing from client...")
			conn.Close()
			return
		}

		responseData := "response"
		responseByteData, err := json.Marshal(responseData)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = conn.Write(responseByteData)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
