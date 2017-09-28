package main


import (
	"fmt"
	"runtime"
	//"time"

)

func main() {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	// number of workers, and size of job queue
	pool := NewPool(65535, 2000000)
	defer pool.Release()

	// how many jobs we should wait
	//pool.WaitCount(3)

	// submit one or more jobs to pool
	for i := 0; i < 2000000; i++ {
		count := i
		pool.WaitCount(1)
		pool.JobQueue <- func() {
			// say that job is done, so we can know how many jobs are finished
			defer pool.JobDone()

			fmt.Printf("hello %d\n", count)
		}
	}

	// wait until we call JobDone for all jobs
	pool.WaitAll()

	//time.Sleep(1 * time.Second)
}

//返回数据 结果 设计思路
//https://gobyexample.com/worker-pools
/*

// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

package main

import "fmt"
import "time"

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {

    // In order to use our pool of workers we need to send
    // them work and collect their results. We make 2
    // channels for this.
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // This starts up 3 workers, initially blocked
    // because there are no jobs yet.
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Here we send 5 `jobs` and then `close` that
    // channel to indicate that's all the work we have.
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // Finally we collect all the results of the work.
    for a := 1; a <= 5; a++ {
        <-results
    }
}
 */

// go http server 已经做了 go
//https://hustcat.github.io/http_server/

