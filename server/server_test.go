package server

import (
	"fmt"
	"testing"
)

func Test_Server(t *testing.T) {

	var num = 0
	listener := NewServer()

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
		}
		go listener.Process(conn)

		num++

		fmt.Println(num)

	}

}
