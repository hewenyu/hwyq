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

		*listener.PChan <- num
		listener.Length += 1
		go listener.Process(conn)

		fmt.Printf("now %v", listener.Length)

	}

}
