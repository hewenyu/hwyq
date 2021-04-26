package server

import (
	"fmt"
	"net"

	"github.com/hewenyu/hwyq/queue"
)

const DefultPort = 8950

/*
ConnectQueue 链接
*/
type ConnectQueue struct {
	Listener *net.TCPListener // 链接
	Ques     *queue.Queue
	PopChan  *chan int
	PushChan *chan int
}

/*
Close 关闭链接
*/
func (c *ConnectQueue) Close() {
	c.Listener.Close()
	close(*c.PopChan)
	close(*c.PushChan)
}

/*
Accept 获取数据
*/
func (c *ConnectQueue) Accept() (conn net.Conn, err error) {

	conn, err = c.Listener.Accept()

	return
}

/*
NewServer 初始化一个server
*/
func NewServer() *ConnectQueue {
	listener, _ := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IPv4zero, Port: DefultPort})

	server := ConnectQueue{
		Listener: listener,
		Ques:     queue.New(),
		PopChan:  MakeChanInt(LimitFork),
		PushChan: MakeChanInt(LimitFork),
	}

	return &server
}

func (c *ConnectQueue) Process(conn net.Conn) {
	defer conn.Close()

	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("Read from tcp server failed,err:", err)
			break
		}
		data := string(buf[:n])
		fmt.Printf("Recived from client,data:%s\n", data)

	}
}
