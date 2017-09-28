package servergo
import (
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	//"strconv"
)
var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "3333", "port")

func main() {
	var wg sync.WaitGroup

	flag.Parse()

	for i:=0; i<40000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn, err := net.Dial("tcp", *host + ":" + *port)
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
		}()
	}
	wg.Wait()
}
func handleWrite(conn net.Conn, done chan string) {
	//for i := 10; i > 0; i-- {
		_, e := conn.Write([]byte("hello,world ! ##################################################@@" +  "\n"))
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

