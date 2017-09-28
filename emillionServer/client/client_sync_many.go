package main
import (
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	//"strconv"
)
var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "8080", "port")

func main() {
	var wg sync.WaitGroup

	flag.Parse()

	for n:=0; n<10; n++ {
		localip := "12.12.12."+ string(n+10)
	for i:=0; i<50000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//conn, err := net.Dial("tcp", *host + ":" + *port)
			localaddr := &net.TCPAddr{IP:net.ParseIP(localip)}
			remoteaddr := &net.TCPAddr{IP:net.ParseIP(*host),Port:8080}
			conn,err := net.DialTCP("tcp4",localaddr,remoteaddr)
			if err != nil {
				fmt.Println("Error connecting:", err)
				os.Exit(1)
			}
			defer conn.Close()
			fmt.Println("Connecting to " + *host + ":" + *port," -> ",(i+1)*(n+1))
			done := make(chan string)
			go handleWrite(conn, done)
			go handleRead(conn, done)
			fmt.Println(<-done)
			fmt.Println(<-done)
		}()
	}
	}
	wg.Wait()
}
func handleWrite(conn net.Conn, done chan string) {
	_, e := conn.Write([]byte("hello,world !" +  "\n"))
	if e != nil {
		fmt.Println("Error to send message because of ", e.Error())
		//break
		os.Exit(1)
	}
	//}
	done <- "Sent"
}
func handleRead(conn net.Conn, done chan string) {
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error to read message because of ", err)
		os.Exit(1)
		return
	}
	fmt.Println(string(buf[:reqLen-1]))
	done <- "Read"
}

