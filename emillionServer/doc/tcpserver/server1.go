package main

import (
	"fmt"
	"net"
	"time"
	"github.com/petermattis/goid"
)

func main() {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
	// 监听并接受来自客户端的连接
	for i:=0 ; i<3; i++ {
		go func(n int) {

			for {

				conn, err := listener.Accept()
				if err != nil {
					fmt.Println("Error accepting", err.Error())
					return // 终止程序
				}
				fmt.Println("goroutine id ---> index   "  , goid.Get(), n)
				go doServerStuff(conn)
			}
		}(i)
	}

	time.Sleep(100000 * time.Second)
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil && len != 0{
			fmt.Println("Error reading", err.Error())
			return //终止程序
		} else if err != nil && len == 0{
			fmt.Println("   client close ", err.Error())
			return //终止程序
		}
		fmt.Printf("Received data: %v", string(buf[:len]))
	}
}
