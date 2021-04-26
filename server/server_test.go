package server

import (
	"fmt"
	"testing"
)

func Test_Server(t *testing.T) {

	listener := NewServer()

	defer listener.Close()
	// var num int

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
		}
		go listener.Process(conn)

	}

}
