// $ 6g echo.go && 6l -o echo echo.6
// $ ./echo
//
//  ~ in another terminal ~
//
// $ nc localhost 3540

package servergo

import (
    "net"
    "bufio"
    "strconv"
    "fmt"

)

const PORT = 3333

func main() {
    server, err := net.Listen("tcp", ":" + strconv.Itoa(PORT))
    if server == nil {
        fmt.Printf("couldn't start listening: %s",err.Error())// + err.String())
           return
    }
    conns := clientConns(server)
    for {
        go handleConn(<-conns)
    }
}

func clientConns(listener net.Listener) chan net.Conn {
    ch := make(chan net.Conn,10000)
    i := 0
    go func() {
        for {
            client, err := listener.Accept()
            if client == nil {
                fmt.Printf("couldn't accept: ",err.Error())// + err.Errors())
                continue
            }
            i++
            fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
            ch <- client
        }
    }()
    return ch
}

func handleConn(client net.Conn) {
    b := bufio.NewReader(client)
    for {
        line, err := b.ReadBytes('\n')
        if err != nil { // EOF, or worse
            break
        }
        client.Write(line)
    }
}


