package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Client() {
	// 连接服务器

	// net.DialTCP("network string", laddr *net.TCPAddr, raddr *net.TCPAddr)
	client, err := net.Dial("tcp", "127.0.0.1:8950")
	if err != nil {
		fmt.Println("Connect to TCP server failed ,err:", err)
		return
	}

	defer client.Close()

	inputReader := bufio.NewReader(os.Stdin)
	// 一直读取直到遇到换行符
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Read from console failed,err:", err)
			return
		}

		// 读取到字符"Q"退出
		str := strings.TrimSpace(input)
		if str == "Q" {
			break
		}

		// 响应服务端信息
		_, err = client.Write([]byte(input))
		if err != nil {
			fmt.Println("Write failed,err:", err)
			break
		}
	}

}
