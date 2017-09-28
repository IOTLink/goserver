package main


import (
"bufio"
"fmt"
"net"
"time"
)

var quitSemaphore chan bool

func main() {
	var tcpAddr *net.TCPAddr
	//localaddr := &net.TCPAddr{IP:net.ParseIP("192.168.1.107")}
	localaddr := &net.TCPAddr{IP:net.ParseIP("192.168.1.107")}
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "192.168.1.106:3333")
	//tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:3333")

	conn, _ := net.DialTCP("tcp", localaddr, tcpAddr)
	//conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")

	go onMessageRecived(conn)

	//b := []byte("time\n")
	ts := time.Now().String()
	b := []byte("frist: "+ts + "\n" )
	conn.Write(b)

	<-quitSemaphore
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
		time.Sleep(time.Second)
		ts := time.Now().String()
		b := []byte("second: "+ts + "\n" )
		conn.Write(b)
	}
}

