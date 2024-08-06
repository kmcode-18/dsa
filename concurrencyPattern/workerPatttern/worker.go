//worker pool pattern
package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("worker %d processing job %d\n", id, job)
		result <- job * 2
	}
}

func main() {
	numJobs := 10
	numWorker := 3
	jobs := make(chan int, numJobs)
	reslut := make(chan int, numJobs)
	var wg sync.WaitGroup
	//workers starting
	for i := 0; i < numWorker; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			worker(workerId, jobs, reslut)
		}(i)
	}
	//enque jobs
	for i := 0; i < numJobs; i++ {
		jobs <- i
	}
	close(jobs)
	//waiting for all goworker to finsh tasks
	go func() {
		wg.Wait()
		close(reslut)
	}()
	//collect tasks
	for r := range reslut {
		fmt.Printf("result:%d\n", r)
	}

}
