package server

import (
	"fmt"
	"io"
	"net"

	"github.com/hewenyu/hwyq/queue"
)

/*
ConnectQueue 链接
*/
type ConnectQueue struct {
	Listener *net.TCPListener // 链接
	Ques     *queue.Queue     // 消息队列的主要链表
	PChan    *chan int        // 限制最高连接数
	Length   int              // 用于统计队列
}

/*
Close 关闭链接
*/
func (c *ConnectQueue) Close() {
	c.Listener.Close()
	close(*c.PChan)
}

/*
Accept 获取数据
*/
func (c *ConnectQueue) Accept() (conn net.Conn, err error) {

	conn, err = c.Listener.Accept()

	return
}

/**
 * NewServer 初始化一个server
 */
func NewServer() *ConnectQueue {
	listener, _ := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IPv4zero, Port: DefultPort})

	server := ConnectQueue{
		Listener: listener,
		Ques:     queue.New(),
		PChan:    MakeChanInt(LimitFork),
		Length:   0,
	}

	return &server
}

func (c *ConnectQueue) Process(conn net.Conn) {
	defer conn.Close()

	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {

			if err != io.EOF {
				fmt.Println("Read from tcp server failed,err:", err)
			}

			break
		}
		data := string(buf[:n])
		fmt.Printf("Recived from client,data:%s\n", data)

		c.Ques.Push(buf[:n])

	}
	fmt.Println("链接结束")
	<-*c.PChan
	c.Length -= 1
}
