package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	//"strconv"
	"github.com/golang/protobuf/proto"
	"../protest"
	"log"
)
var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "3333", "port")
func main() {
	flag.Parse()


	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to " + *host + ":" + *port)
	done := make(chan string)
	go handleWrite(conn, done)
	go handleRead(conn, done)
	fmt.Println(<-done)
	fmt.Println(<-done)
}
func handleWrite(conn net.Conn, done chan string) {
	test := &protest.Test {
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
	}
	fmt.Println(test)

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	_, e := conn.Write(data)
	if e != nil {
		fmt.Println("Error to send message because of ", e.Error())
		//break
	}
	//}
	done <- "Sent"
}
func handleRead(conn net.Conn, done chan string) {
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error to read message because of ", err)
		return
	}
	//fmt.Println(string(buf[:reqLen-1]))
	protodata := new(protest.Test)
	//Convert all the data retrieved into the ProtobufTest.TestMessage struct type
	err = proto.Unmarshal(buf[0:reqLen], protodata)
	fmt.Println(reqLen," --- ",protodata)
	done <- "Read"
}

