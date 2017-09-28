
package main

import (

"fmt"
"net"
"sync"
	"time"
)

func main() {
	//打开连接:
	var wg sync.WaitGroup

	t1 := time.Now() // get current time

	for i:= 0; i<30000; i++ {
		wg.Add(1)
		go func() {
			conn, err := net.Dial("tcp", "127.0.0.1:50000")
			if err != nil {
				//由于目标计算机积极拒绝而无法创建连接
				fmt.Println("Error dialing", err.Error())
				wg.Done()
				return // 终止程序
			}

			_, err = conn.Write([]byte("lhy says: hello world"))
			wg.Done()
		}()
	}

	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)

	wg.Wait()
	return

}

