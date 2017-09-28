// $ 6g echo.go && 6l -o echo echo.6
// $ ./echo
//
//  ~ in another terminal ~
//
// $ nc localhost 3540

package main

import (
    "net"
    "strconv"
    "fmt"
    "runtime"
    "os"
    "os/signal"
    . "emillionServer/toolbox"
    "time"
)

const PORT = 8080

func WatchProcess() {
    go func() {
        //ProcessInput("lookup goroutine", os.Stdout)
        //ProcessInput("lookup heap", os.Stdout)
        //ProcessInput("lookup threadcreate", os.Stdout)
        //ProcessInput("lookup block", os.Stdout)
        for {
            time.Sleep(time.Second * 1)
            ProcessInput("gc summary", os.Stdout)
        }
    }()
}

func main() {
    i := 0
    numCPUs := runtime.NumCPU()
    runtime.GOMAXPROCS(numCPUs)

    // number of workers, and size of job queue
    pool := NewPool(60000, 2000000)
    defer pool.Release()

    fmt.Printf("init pool success!\n")
    server, err := net.Listen("tcp", ":" + strconv.Itoa(PORT))
    defer server.Close()

    if server == nil {
        fmt.Printf("couldn't start listening: %s",err.Error())// + err.String())
        return
    }

    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, os.Interrupt, os.Kill)
    go func() {
        for sig := range sigs {
            fmt.Printf("received sigs %v\n", sig)
            pool.Release()
            server.Close()
            os.Exit(-1)
        }
    }()
    //WatchProcess()

    for {
        client, err := server.Accept()
        if client == nil {
            fmt.Printf("couldn't accept: ",err.Error())// + err.Errors())
            continue
        }
        i++
        fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
        pool.JobQueue <- client
    }

}
