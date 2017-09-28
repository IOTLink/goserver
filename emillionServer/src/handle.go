package main

import (
	"net"
	"bufio"
	//"fmt"
	"time"
	"fmt"
)


func handleConn(client net.Conn) {
	//ipStr := client.RemoteAddr().String()
	/*
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		client.Close()
	}()
	*/
	defer client.Close()
	reader := bufio.NewReader(client)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("client close....")
			return
		}

		//fmt.Println(string(message))
		msg := time.Now().String() + message + "\n"
		b := []byte(msg)
		client.Write(b)
		//client.Close()
	}
}
