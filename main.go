package main

import "github.com/hewenyu/hwyq/client"

func main() {

	// var num = 0
	// listener := server.NewServer()

	// defer listener.Close()

	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		fmt.Println("accept err:", err)
	// 	}

	// 	*listener.PushChan <- num
	// 	listener.PushLength += 1
	// 	go listener.Process(conn)

	// 	fmt.Printf("now %v\n", listener.PushLength)
	// }

	client.Client()

}
